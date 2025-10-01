package main

import (
	"backend/config"
	"backend/controller"
	"backend/middleware"
	"backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	config.InitDB()
	controller.InitTodoController(config.DB)

	r := gin.Default()
	r.RedirectTrailingSlash = false
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://26.1.224.212:3000", "http://26.176.162.130:3000", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	routes.TodoRoutes(r)
	routes.UserRoutes(r)

	jwtKey := []byte(os.Getenv("JWT_KEY"))
	port := os.Getenv("PORT")

	middleware.InitAuth(jwtKey)
	if port == "" {
		port = "8080"
	}
	log.Println("🚀 Сервер запущен на http://localhost:" + port)
	r.Run(":" + port)
}
