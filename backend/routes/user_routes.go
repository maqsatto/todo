package routes

import (
	"backend/controller"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	user := r.Group("/users")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
	}

	protected := r.Group("/profile")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/", func(c *gin.Context) {
			email := c.GetString("email")
			c.JSON(200, gin.H{"message": "Hello " + email})
		})

		protected.PUT("/username", controller.UpdateUsername)
		protected.PUT("/password", controller.UpdatePassword)
		protected.PUT("/pfp", controller.UpdatePFP)
	}
}
