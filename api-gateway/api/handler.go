package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/NeverlandMJ/ToDo/api-gateway/pkg/auth"
	"github.com/NeverlandMJ/ToDo/api-gateway/pkg/entity"
	customErr "github.com/NeverlandMJ/ToDo/api-gateway/pkg/error"
	"github.com/NeverlandMJ/ToDo/api-gateway/service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	_ "github.com/NeverlandMJ/ToDo/api-gateway/docs"
)

type Handler struct {
	provider service.Provider
}

func NewHandler(prov service.Provider) Handler {
	return Handler{
		provider: prov,
	}
}

type message struct {
	Message interface{} `json:"message"`
	Success bool        `json:"success"`
}

// ShowAccount godoc
// @Summary      Send TOTP
// @Description  send one time code to user's phone
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param 		 request body entity.ReqPhone true "user's phone number"
// @Success      200  {object}  message
// @Failure      404  {object}  message
// @Failure      500  {object}  message
// @Router        /auth/v1/userpb/send-code [POST]
func (h Handler) SendCode(c *gin.Context) {
	var ph entity.ReqPhone
	if err := c.BindJSON(&ph); err != nil {
		r := message{
			Message: "invalid json",
			Success: false,
		}
		c.JSON(http.StatusBadRequest, r)
		fmt.Println(err)
		return
	}

	if err := ph.CheckReqPhone(); err != nil {
		h.HandleErr(err, c)
		return
	}

	resp, err := h.provider.UserServiceProvider.SendCode(c.Request.Context(), ph)
	if err != nil {
		h.HandleErr(err, c)
		return
	}
	r := message{
		Message: fmt.Sprintf("code has been sent to %s", resp.Phone),
		Success: true,
	}
	c.JSON(http.StatusOK, r)
}

// ShowAccount godoc
// @Summary      Sign up
// @Description  sign up with TOTP
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body entity.ReqSignUp true "user's phone number and one time code which was sent to user's phone"
// @Success      200  {object}  entity.RespUser
// @Failure      404  {object}  message
// @Failure      409  {object}  message
// @Failure      408  {object}  message
// @Failure      401  {object}  message
// @Failure      500  {object}  message
// @Router        /auth/v1/userpb/register [POST]
func (h Handler) SignUp(c *gin.Context) {
	var cd entity.ReqSignUp
	if err := c.BindJSON(&cd); err != nil {
		r := message{
			Message: "invalid json",
			Success: false,
		}
		c.JSON(http.StatusBadRequest, r)
		fmt.Println(err)
		return
	}

	if err := cd.CheckReqCode(); err != nil {
		h.HandleErr(err, c)
		return
	}

	resp, err := h.provider.UserServiceProvider.RegisterUser(c.Request.Context(), cd)
	if err != nil {
		h.HandleErr(err, c)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// ShowAccount godoc
// @Summary      Sign in
// @Description  sign in with default user_name and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param 		 request body entity.ReqSignIn true "user_name and password"
// @Success      200  {object}  message
// @Failure      404  {object}  message
// @Failure      409  {object}  message
// @Failure      401  {object}  message
// @Failure      500  {object}  message
// @Router        /auth/v1/userpb/sign-in [POST]
func (h Handler) SignIn(c *gin.Context) {
	var data entity.ReqSignIn

	if err := c.BindJSON(&data); err != nil {
		r := message{
			Message: "invalid json",
			Success: false,
		}
		c.JSON(http.StatusBadRequest, r)
		fmt.Println(err)
		return
	}

	if err := data.CheckReqSignIn(); err != nil {
		h.HandleErr(err, c)
		return
	}

	token, err := h.provider.UserServiceProvider.SignIn(context.Background(), data)
	if err != nil {
		h.HandleErr(err, c)
		return
	}

	c.SetCookie(
		"token",
		token,
		3600,
		"/",
		"localhost",
		false,
		true,
	)

	r := message{
		Message: "succesfully loged in",
		Success: true,
	}
	c.JSON(http.StatusOK, r)
}

// ShowAccount godoc
// @Summary      Change password
// @Description  change user's password
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 		 request body entity.ReqChangePassword true "old and new password"
// @Success      200  {object}  message
// @Failure      404  {object}  message
// @Failure      409  {object}  message
// @Failure      401  {object}  message
// @Failure      403  {object}  message
// @Failure      500  {object}  message
// @Router        /api/v1/userpb/change/password/ [PUT]
func (h Handler) ChangePassword(c *gin.Context)  {
	v, ok := c.Get("claims")
	if !ok {
		return
	}
	claims, ok := v.(*auth.Claims)
	if !ok {
		r := message{
			Message: "looks like cookie isn't set",
			Success: false,
		}
		c.JSON(http.StatusUnauthorized, r)
		return
	}

	var r entity.ReqChangePassword
	if err := c.BindJSON(&r); err != nil {
		r := message{
			Message: "invalid json",
			Success: false,
		}
		c.JSON(http.StatusBadRequest, r)
		fmt.Println(err)
		return
	}

	if err := r.CheckReqChangePassword(); err != nil {
		h.HandleErr(err, c)
		return
	}

	err := h.provider.UserServiceProvider.ChangePassword(context.Background(), claims.ID, r)
	if err != nil {
		h.HandleErr(err, c)
		return
	}

	m := message{
		Message: "password succesfully changed",
		Success: true,
	}
	c.JSON(http.StatusOK, m)

}

// ShowAccount godoc
// @Summary      Change user_name
// @Description  change user's user_name
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 		 request body entity.ReqChangeUsername true "new user_name"
// @Success      200  {object}  message
// @Failure      404  {object}  message
// @Failure      409  {object}  message
// @Failure      401  {object}  message
// @Failure      403  {object}  message
// @Failure      500  {object}  message
// @Router        /api/v1/userpb/change/user-name/ [PUT]
func (h Handler) ChangeUserName(c *gin.Context)  {
	v, ok := c.Get("claims")
	if !ok {
		return
	}
	claims, ok := v.(*auth.Claims)
	if !ok {
		r := message{
			Message: "looks like cookie isn't set",
			Success: false,
		}
		c.JSON(http.StatusUnauthorized, r)
		return
	}

	var r entity.ReqChangeUsername
	if err := c.BindJSON(&r); err != nil {
		r := message{
			Message: "invalid json",
			Success: false,
		}
		c.JSON(http.StatusBadRequest, r)
		fmt.Println(err)
		return
	}

	if err := r.CheckReqChangeUsername(); err != nil {
		h.HandleErr(err, c)
		return
	}

	err := h.provider.UserServiceProvider.ChangeUserName(context.Background(), claims.ID, r)
	if err != nil {
		h.HandleErr(err, c)
		return
	}

	m := message{
		Message: "user name succesfully changed",
		Success: true,
	}
	c.JSON(http.StatusOK, m)
}

// ShowAccount godoc
// @Summary      delete account
// @Description  delete current account
// @Tags         user
// @Accept       json
// @Produce      json
// @Param 		 request body entity.ReqSignIn true "user signs in"
// @Success      200  {object}  message
// @Failure      404  {object}  message
// @Failure      409  {object}  message
// @Failure      401  {object}  message
// @Failure      403  {object}  message
// @Failure      500  {object}  message
// @Router        /api/v1/userpb/delete/account/ [DELETE]
func (h Handler) DeleteAccount(c *gin.Context)  {
	v, ok := c.Get("claims")
	if !ok {
		return
	}
	claims, ok := v.(*auth.Claims)
	if !ok {
		r := message{
			Message: "looks like cookie isn't set",
			Success: false,
		}
		c.JSON(http.StatusUnauthorized, r)
		return
	}

	var r entity.ReqSignIn
	if err := c.BindJSON(&r); err != nil {
		r := message{
			Message: "invalid json",
			Success: false,
		}
		c.JSON(http.StatusBadRequest, r)
		fmt.Println(err)
		return
	}

	if err := r.CheckReqSignIn(); err != nil {
		h.HandleErr(err, c)
		return
	}

	err := h.provider.UserServiceProvider.DeleteAccount(context.Background(), claims.Id, r)
	if err != nil {
		h.HandleErr(err, c)
		return
	}

	m := message{
		Message: "accaunt deleted",
		Success: true,
	}
	c.JSON(http.StatusOK, m)
}

// ShowAccount godoc
// @Summary      Create todo
// @Description  create todo
// @Tags         todo
// @Accept       json
// @Produce      json
// @Param 		 request body entity.ReqCreateTodo true "todo's body"
// @Success      200  {object}  entity.RespTodo
// @Failure      404  {object}  message
// @Failure      409  {object}  message
// @Failure      401  {object}  message
// @Failure      500  {object}  message
// @Router        /api/v1/todopb/create [POST]
func (h Handler) CreateTodo(c *gin.Context) {
	v, ok := c.Get("claims")
	if !ok {
		return
	}
	claims, ok := v.(*auth.Claims)
	if !ok {
		r := message{
			Message: "looks like cookie isn't set",
			Success: false,
		}
		c.JSON(http.StatusUnauthorized, r)
		return
	}

	var td entity.ReqCreateTodo
	if err := c.BindJSON(&td); err != nil {
		r := message{
			Message: "invalid json",
			Success: false,
		}
		c.JSON(http.StatusBadRequest, r)
		fmt.Println(err)
		return
	}

	if err := td.CheckReqCreateTodo(); err != nil {
		h.HandleErr(err, c)
		return
	}

	new, err := h.provider.TodoServiceProvider.CreateTodo(c.Request.Context(), td, claims.ID)
	if err != nil {
		h.HandleErr(err, c)
		return
	}

	c.JSON(http.StatusCreated, new)
}

// ShowAccount godoc
// @Summary      Get todo
// @Description  get todo by id
// @Tags         todo
// @Produce      json
// @Param        todo-id   path      string  true  "Todo ID"
// @Success      200  {object}  entity.RespTodo
// @Failure      404  {object}  message
// @Failure      401  {object}  message
// @Failure      500  {object}  message
// @Router        /api/v1/todopb/get/{todo-id} [GET]
func (h Handler) GetTodoByID(c *gin.Context) {
	v, ok := c.Get("claims")
	if !ok {
		return
	}
	claims, ok := v.(*auth.Claims)
	// fmt.Println(claims)
	if !ok {
		r := message{
			Message: "looks like cookie isn't set",
			Success: false,
		}
		c.JSON(http.StatusUnauthorized, r)
		return
	}

	todoID, ok := c.Params.Get("todo-id")
	if !ok {
		r := message{
			Message: "invalid params",
			Success: false,
		}
		c.JSON(http.StatusBadRequest, r)
	}

	resp, err := h.provider.TodoServiceProvider.GetTodoByID(context.Background(), claims.ID, todoID)
	if err != nil {
		h.HandleErr(err, c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ShowAccount godoc
// @Summary      Mark as done
// @Description  Mark todo as done by todo's ID
// @Tags         todo
// @Produce      json
// @Param        todo-id   path      string  true  "Todo ID"
// @Success      200  {object}  message
// @Failure      404  {object}  message
// @Failure      401  {object}  message
// @Failure      500  {object}  message
// @Router       /api/v1/todopb/done/{todo-id} [PUT]
func (h Handler) MarkAsDone(c *gin.Context) {
	v, ok := c.Get("claims")
	if !ok {
		return
	}
	claims, ok := v.(*auth.Claims)
	// fmt.Println(claims)
	if !ok {
		r := message{
			Message: "looks like cookie isn't set",
			Success: false,
		}
		c.JSON(http.StatusUnauthorized, r)
		return
	}

	todoID, ok := c.Params.Get("todo-id")
	if !ok {
		r := message{
			Message: "invalid params",
			Success: false,
		}
		c.JSON(http.StatusBadRequest, r)
	}

	err := h.provider.TodoServiceProvider.MarkAsDone(context.Background(), claims.ID, todoID)
	if err != nil {
		h.HandleErr(err, c)
		return
	}

	r := message{
		Message: "successfully updated",
		Success: true,
	}
	c.JSON(http.StatusOK, r)
}

// ShowAccount godoc
// @Summary      Delete todo
// @Description  Delete todo by todo's ID
// @Tags         todo
// @Produce      json
// @Param        todo-id   path      string  true  "Todo ID"
// @Success      200  {object}  message
// @Failure      404  {object}  message
// @Failure      401  {object}  message
// @Failure      500  {object}  message
// @Router       /api/v1/todopb/delete/{todo-id} [DELETE]
func (h Handler) DeleteTodoByID(c *gin.Context) {
	v, ok := c.Get("claims")
	if !ok {
		return
	}
	claims, ok := v.(*auth.Claims)
	// fmt.Println(claims)
	if !ok {
		r := message{
			Message: "looks like cookie isn't set",
			Success: false,
		}
		c.JSON(http.StatusUnauthorized, r)
		return
	}

	todoID, ok := c.Params.Get("todo-id")
	if !ok {
		r := message{
			Message: "invalid params",
			Success: false,
		}
		c.JSON(http.StatusBadRequest, r)
	}

	err := h.provider.TodoServiceProvider.DeleteTodoByID(context.Background(), claims.ID, todoID)
	if err != nil {
		h.HandleErr(err, c)
		return
	}

	r := message{
		Message: "successfully deleted",
		Success: true,
	}
	c.JSON(http.StatusOK, r)
}

// ShowAccount godoc
// @Summary      Get all todos
// @Description  Get all todo by userID
// @Tags         todo
// @Produce      json
// @Success      200  {object}  []entity.RespTodo
// @Failure      404  {object}  message
// @Failure      401  {object}  message
// @Failure      500  {object}  message
// @Router       /api/v1/todopb/get/todos [GET]
func (h Handler) GetAllTodos(c *gin.Context) {
	v, ok := c.Get("claims")
	if !ok {
		return
	}
	claims, ok := v.(*auth.Claims)
	// fmt.Println(claims)
	if !ok {
		r := message{
			Message: "looks like cookie isn't set",
			Success: false,
		}
		c.JSON(http.StatusUnauthorized, r)
		return
	}

	resp, err := h.provider.TodoServiceProvider.GetAllTodos(context.Background(), claims.ID)
	if err != nil {
		h.HandleErr(err, c)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ShowAccount godoc
// @Summary      Update todo's body
// @Description  Update todo's body
// @Tags         todo
// @Produce      json
// @Param 		 request body entity.ReqUpdateBody true "todo's body"
// @Success      200  {object}  message
// @Failure      404  {object}  message
// @Failure      401  {object}  message
// @Failure      500  {object}  message
// @Router       /api/v1/todopb/update/body/ [PUT]
func (h Handler) UpdateTodosBody(c *gin.Context) {
	v, ok := c.Get("claims")
	if !ok {
		return
	}
	claims, ok := v.(*auth.Claims)
	if !ok {
		r := message{
			Message: "looks like cookie isn't set",
			Success: false,
		}
		c.JSON(http.StatusUnauthorized, r)
		return
	}

	var r entity.ReqUpdateBody
	if err := c.BindJSON(&r); err != nil {
		r := message{
			Message: "invalid json",
			Success: false,
		}
		c.JSON(http.StatusBadRequest, r)
		fmt.Println(err)
		return
	}

	if err := r.CheckReqUpdateBody(); err != nil {
		h.HandleErr(err, c)
		return
	}

	err := h.provider.TodoServiceProvider.UpdateTodosBody(context.Background(), claims.ID, r)
	if err != nil {
		h.HandleErr(err, c)
		return
	}

	m := message{
		Message: "succesfully updated",
		Success: true,
	}
	c.JSON(http.StatusOK, m)
}

// ShowAccount godoc
// @Summary      Update todo's deadline
// @Description  Update todo's deadline
// @Tags         todo
// @Produce      json
// @Param 		 request body entity.ReqUpdateDeadline true "todo's body"
// @Success      200  {object}  message
// @Failure      404  {object}  message
// @Failure      401  {object}  message
// @Failure      500  {object}  message
// @Router       /api/v1/todopb/update/deadline/ [PUT]
func (h Handler) UpdateTodosDeadline(c *gin.Context) {
	v, ok := c.Get("claims")
	if !ok {
		return
	}
	claims, ok := v.(*auth.Claims)
	if !ok {
		r := message{
			Message: "looks like cookie isn't set",
			Success: false,
		}
		c.JSON(http.StatusUnauthorized, r)
		return
	}
	var r entity.ReqUpdateDeadline
	if err := c.BindJSON(&r); err != nil {
		r := message{
			Message: "invalid json",
			Success: false,
		}
		c.JSON(http.StatusBadRequest, r)
		fmt.Println(err)
		return
	}

	if err := r.CheckReqUpdateDeadline(); err != nil {
		h.HandleErr(err, c)
		return
	}

	err := h.provider.TodoServiceProvider.UpdateTodosDeadline(context.Background(), claims.ID, r)
	if err != nil {
		h.HandleErr(err, c)
		return
	}

	m := message{
		Message: "succesfully updated",
		Success: true,
	}
	c.JSON(http.StatusOK, m)
}

// ShowAccount godoc
// @Summary      Delete done
// @Description  Delete todos which were marked as done
// @Tags         todo
// @Produce      json
// @Success      200  {object}  message
// @Failure      404  {object}  message
// @Failure      401  {object}  message
// @Failure      500  {object}  message
// @Router       /api/v1/todopb/delete/done [DELETE]
func (h Handler) DeleteDoneTodos(c *gin.Context) {
	v, ok := c.Get("claims")
	if !ok {
		return
	}
	claims, ok := v.(*auth.Claims)
	if !ok {
		r := message{
			Message: "looks like cookie isn't set",
			Success: false,
		}
		c.JSON(http.StatusUnauthorized, r)
		return
	}

	err := h.provider.TodoServiceProvider.DeleteDoneTodos(context.Background(), claims.ID)
	if err != nil {
		h.HandleErr(err, c)
		return
	}

	m := message{
		Message: "succesfully deleted",
		Success: true,
	}
	c.JSON(http.StatusOK, m)
}

// ShowAccount godoc
// @Summary      Delete passed deadline
// @Description  Delete todos which's deadline had already passed
// @Tags         todo
// @Produce      json
// @Success      200  {object}  message
// @Failure      404  {object}  message
// @Failure      401  {object}  message
// @Failure      500  {object}  message
// @Router       /api/v1/todopb/delete/passed [DELETE]
func (h Handler) DeletePassedDeadline(c *gin.Context)  {
	v, ok := c.Get("claims")
	if !ok {
		return
	}
	claims, ok := v.(*auth.Claims)
	if !ok {
		r := message{
			Message: "looks like cookie isn't set",
			Success: false,
		}
		c.JSON(http.StatusUnauthorized, r)
		return
	}

	err := h.provider.TodoServiceProvider.DeletePassedDeadline(context.Background(), claims.ID)
	if err != nil {
		h.HandleErr(err, c)
		return
	}

	m := message{
		Message: "succesfully deleted",
		Success: true,
	}
	c.JSON(http.StatusOK, m)
}

func (h Handler) HandleErr(err error, c *gin.Context) {
	log.Println(err)
	if sts, ok := status.FromError(err); ok {
		switch sts.Code() {
		case codes.Internal:
			r := message{
				Message: "internal server error occured",
				Success: false,
			}
			c.JSON(http.StatusInternalServerError, r)
		case codes.AlreadyExists:
			r := message{
				Message: "this user is already exist",
				Success: false,
			}
			c.JSON(http.StatusConflict, r)
		case codes.NotFound:
			r := message{
				Message: sts.Message(),
				Success: false,
			}
			c.JSON(http.StatusNotFound, r)
		case codes.Unauthenticated:
			r := message{
				Message: sts.Message(),
				Success: false,
			}
			c.JSON(http.StatusUnauthorized, r)
		case codes.InvalidArgument:
			r := message{
				Message: sts.Message(),
				Success: false,
			}
			c.JSON(http.StatusBadRequest, r)
		case codes.PermissionDenied:
			r := message{
				Message: sts.Message(),
				Success: false,
			}
			c.JSON(http.StatusForbidden, r)
		default:
			r := message{
				Message: sts.Message(),
				Success: false,
			}
			c.JSON(http.StatusInternalServerError, r)
		}
	} else if errors.Is(err, customErr.ERR_CODE_HAS_EXPIRED) {
		r := message{
			Message: "code has expired",
			Success: false,
		}
		c.JSON(http.StatusRequestTimeout, r)
	} else if errors.Is(err, customErr.ERR_USER_BLOCKED) {
		r := message{
			Message: "current user is blocked",
			Success: false,
		}
		c.JSON(http.StatusForbidden, r)
	} else if errors.Is(err, customErr.ERR_INVALID_INPUT) {
		r := message{
			Message: err.Error(),
			Success: false,
		}
		c.JSON(http.StatusBadRequest, r)
	} else {
		r := message{
			Message: "unexpected server error",
			Success: false,
		}
		c.JSON(http.StatusInternalServerError, r)
	}

}
