package middlewares

import (
	"net/http"

	"github.com/NeverlandMJ/ToDo/api-gateway/pkg/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type message struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success,omitempty"`
}

// Authentication middleware for authentication
func Authentication(c *gin.Context) {
	cook, err := c.Request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			r := message{
				Message: "user is not signed in",
				Success: false,
			}
			c.JSON(http.StatusUnauthorized, r)
			return
		} else {
			r := message{
				Message: "invalid request has been made",
				Success: false,
			}
			c.JSON(http.StatusBadRequest, r)
		}
		return
	}

	tokenStr := cook.Value
	claims := &auth.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return auth.JwtKey, nil
	})

	if err != nil {
		if err == http.ErrNoCookie {
			r := message{
				Message: "user is not registrated",
				Success: false,
			}
			c.JSON(http.StatusUnauthorized, r)
			return
		}else {
			r := message{
				Message: "invalid request has been made",
				Success: false,
			}
			c.JSON(http.StatusBadRequest, r)
		}
		return
	}
	if !tkn.Valid {
		r := message{
			Message: "token has been expired please login again",
			Success: false,
		}
		c.JSON(http.StatusUnauthorized, r)
		return
	}

	c.Set("claims", claims)
	
	c.Next()
}

