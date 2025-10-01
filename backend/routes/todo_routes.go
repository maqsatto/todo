package routes

import (
	"backend/controller"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(r *gin.Engine) {
	todo := r.Group("/todos")
	todo.Use(middleware.AuthMiddleware())
	{
		todo.POST("/", controller.CreateTodo)
		todo.GET("/", controller.GetTodos)
		todo.PUT("/:id", controller.UpdateTodo)
		todo.DELETE("/:id", controller.DeleteTodo)
	}
}
