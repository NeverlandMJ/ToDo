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
	Success bool   `json:"success"`
}

// sending code to the user
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

func (h Handler) SignUp(c *gin.Context) {
	var cd entity.ReqCode
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


func (h Handler) CreateTodo(c *gin.Context) {
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

	r := message{
		Message: resp,
		Success: true,
	}
	c.JSON(http.StatusOK, r)
}

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
