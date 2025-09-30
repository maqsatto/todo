package controller

import (
	"backend/models"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var JwtKey []byte

func InitUserController(db *gorm.DB, key []byte) {
	DB = db
	JwtKey = key
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(user.Password) < 6 {
		c.JSON(500, gin.H{"error": "Password must contain at least 6 characters!"})
		return
	}
	var existingUser models.User
	if err := DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	token, err := utils.GenerateJWT(user.Email, JwtKey)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not generate token"})
		return
	}
	DB.Create(&user)
	c.JSON(200, gin.H{
		"token": token,
		"user":  user,
	})
}
func Login(c *gin.Context) {
	var user models.User
	var UserInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&UserInput); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result := DB.Where("email = ?", UserInput.Email).First(&user)
	if result.Error != nil {
		c.JSON(400, gin.H{"error": "Something not matching"})
		return
	}
	if user.Password != UserInput.Password {
		c.JSON(401, gin.H{"error": "Something not matching"})
		return
	}
	token, err := utils.GenerateJWT(user.Email, JwtKey)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
		"user":  user,
	})

}
