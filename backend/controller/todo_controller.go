package controller

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitTodoController(db *gorm.DB) {
	DB = db
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&todo)
	c.JSON(http.StatusOK, todo)
}

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
