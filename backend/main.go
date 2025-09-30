package main

import (
	"backend/controller"
	"backend/models"
	"backend/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var JwtKey = []byte("BAUKA_GOI")

func main() {
	// Настройки подключения к БД
	dsn := "host=localhost user=postgres password=0000 dbname=todolist port=5432 sslmode=disable"

	// Подключение к БД
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Ошибка подключения к БД: ", err)
	}

	// Миграции (создание таблиц)
	err = db.AutoMigrate(&models.User{}, &models.Todo{})
	if err != nil {
		log.Fatal("❌ Ошибка миграции: ", err)
	}

	// Инициализация контроллеров (даем доступ к БД)
	controller.InitTodoController(db)

	// Инициализация роутера
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://26.1.224.212:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Тестовый эндпоинт
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Роуты для ToDo
	routes.TodoRoutes(r)
	routes.UserRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("🚀 Сервер запущен на http://localhost:" + port)
	r.Run(":" + port)
}
