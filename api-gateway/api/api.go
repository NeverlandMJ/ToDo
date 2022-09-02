package api

import (
	"github.com/NeverlandMJ/ToDo/api-gateway/pkg/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(h Handler) *gin.Engine{
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	auth := router.Group("/api/auth")
	auth.POST("/v1/userpb/send-code", h.SendCode)
	auth.POST("/v1/userpb/register", h.SignUp)
	auth.POST("/v1/userpb/sign-in", h.SignIn)
	auth.DELETE("/logout",)

	authored := router.Group("/api")
	authored.Use(middlewares.Authentication)
	authored.POST("/v1/todopb/create", h.CreateTodo)
	authored.GET("/v1/todopb/get/:todo-id", h.GetTodoByID)
	authored.PUT("/v1/todopb/done/:todo-id", h.MarkAsDone)


	return router
}