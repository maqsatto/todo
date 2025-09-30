package routes

import (
	"backend/controller"

	"github.com/gin-gonic/gin"
)

func TodoRoutes(r *gin.Engine) {
	todo := r.Group("/todos")
	{
		todo.POST("/", controller.CreateTodo)
		todo.GET("/", controller.GetTodos)
		todo.PUT("/:id", controller.UpdateTodo)
		todo.DELETE("/:id", controller.DeleteTodo)
	}
}
