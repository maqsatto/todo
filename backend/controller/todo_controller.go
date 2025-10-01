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

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	if todo.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo title cannot be empty"})
		return
	}

	todo.UserID = userID.(uint)

	DB.Create(&todo)
	c.JSON(http.StatusOK, todo)
}

func GetTodos(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var todos []models.Todo

	if err := DB.Where("user_id = ?", userID.(uint)).Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch todos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todos fetch successfully",
		"todos":   todos,
		"success": true,
	})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	if err := DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists || todo.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this todo"})
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

	userID, exists := c.Get("user_id")
	if !exists || todo.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this todo"})
		return
	}

	DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
