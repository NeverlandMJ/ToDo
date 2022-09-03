package api

import (
	"github.com/NeverlandMJ/ToDo/api-gateway/pkg/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
        ginSwagger "github.com/swaggo/gin-swagger"
        
        _ "github.com/NeverlandMJ/ToDo/api-gateway/docs"
)

// @title Go + Gin Todo API
// @version 1.0
// @description This is a sample server todo server. You can visit the GitHub repository at https://github.com/NeverlandMJ/ToDo
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func NewRouter(h Handler) *gin.Engine{
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	auth.POST("/v1/userpb/send-code", h.SendCode)
	auth.POST("/v1/userpb/register", h.SignUp)
	auth.POST("/v1/userpb/sign-in", h.SignIn)
	auth.DELETE("/logout",)

	authored := router.Group("/api")
	authored.Use(middlewares.Authentication)
	authored.POST("/v1/todopb/create", h.CreateTodo)
	authored.GET("/v1/todopb/get/:todo-id", h.GetTodoByID)
	authored.PUT("/v1/todopb/done/:todo-id", h.MarkAsDone)
	authored.DELETE("/v1/todopb/delete/:todo-id", h.DeleteTodoByID)
	authored.GET("/v1/todopb/get/todos", h.GetAllTodos)
	authored.PUT("/v1/todopb/update/body/", h.UpdateTodosBody)
	authored.PUT("/v1/todopb/update/deadline/", h.UpdateTodosDeadline)
	authored.DELETE("/v1/todopb/delete/done/", h.DeleteDoneTodos)
	authored.DELETE("/v1/todopb/delete/passed/", h.DeletePassedDeadline)

	authored.PUT("/v1/userpb/change/password/", h.ChangePassword)
	authored.PUT("/v1/userpb/change/user-name/", h.ChangeUserName)
	authored.DELETE("/v1/userpb/delete/account/", h.DeleteAccount)

	return router
}