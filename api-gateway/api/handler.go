package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NeverlandMJ/ToDo/api-gateway/pkg/entity"
	"github.com/NeverlandMJ/ToDo/api-gateway/service"
	"github.com/gin-gonic/gin"
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
	Message string `json:"message"`
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

	resp, err := h.provider.UserServiceProvider.SendCode(c.Request.Context(), ph)
	if err != nil {
		log.Println(err)
		r := message{
			Message: "error in sending code",
			Success: false,
		}
		c.JSON(http.StatusInternalServerError, r)
		fmt.Println(err)
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

	resp, err := h.provider.UserServiceProvider.RegisterUser(c.Request.Context(), cd)
	if err != nil {
		if err == service.ERR_CODE_HAS_EXPIRED {
			r := message{
				Message: "code has expired",
				Success: false,
			}
			c.JSON(http.StatusBadRequest, r)
			return
		}
		r := message{
			Message: "code doesn't match",
			Success: false,
		}
		c.JSON(http.StatusInternalServerError, r)
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
