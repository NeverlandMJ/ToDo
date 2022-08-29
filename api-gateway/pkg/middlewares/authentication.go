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

func Authentication(c *gin.Context) {
	cook, err := c.Request.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			r := message{
				Message: "user ro'yxatdan o'tmagan",
				Success: false,
			}
			c.JSON(http.StatusUnauthorized, r)
			return
		}
		r := message{
			Message: "noto'g'ri so'rov amalga oshirildi",
			Success: false,
		}
		c.JSON(http.StatusBadRequest, r)
		return
	}

	tokenStr := cook.Value
	claims := &auth.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return auth.JwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			r := message{
				Message: "user ro'yxatdan o'tmagan",
				Success: false,
			}
			c.JSON(http.StatusUnauthorized, r)
			return
		}
		r := message{
			Message: "noto'g'ri so'rov amalga oshirildi",
			Success: false,
		}
		c.JSON(http.StatusBadRequest, r)
		return
	}
	if !tkn.Valid {
		r := message{
			Message: "user ro'yxatdan o'tmagan yoki tokenni vaqti o'tib ketgan",
			Success: false,
		}
		c.JSON(http.StatusUnauthorized, r)
		return
	}

	c.Set("claims", claims)

	c.Next()
}
