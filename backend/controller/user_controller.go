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
		c.JSON(400, gin.H{"error": "Password must contain at least 6 characters"})
		return
	}
	if len(user.Username) < 3 {
		c.JSON(400, gin.H{"error": "Username must be at least 3 characters"})
		return
	}

	var existingUser models.User
	if err := DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	if err := DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Could not create user"})
		return
	}

	token, err := utils.GenerateJWT(user, JwtKey)
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
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Something is wrong"})
		return
	}

	if user.Password != input.Password {
		c.JSON(401, gin.H{"error": "Something is wrong"})
		return
	}

	token, err := utils.GenerateJWT(user, JwtKey)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
		"user":  user,
	})
}

func UpdateUsername(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var input struct {
		Username string `json:"username"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(input.Username) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username must be at least 3 characters"})
		return
	}

	var user models.User
	if err := DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Username = input.Username
	DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "Username updated", "user": user})
}

func UpdatePassword(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var input struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(input.NewPassword) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New password must be at least 6 characters"})
		return
	}

	var user models.User
	if err := DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.Password != input.OldPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Old password is incorrect"})
		return
	}

	if user.Password == input.NewPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New password cannot be the same as the old one"})
		return
	}

	user.Password = input.NewPassword
	DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

func UpdatePFP(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var input struct {
		Pfp string `json:"img"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Pfp = input.Pfp
	DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "Profile picture updated", "user": user})
}
