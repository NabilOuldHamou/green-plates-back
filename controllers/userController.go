package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"green-plates-back/initializers"
	"green-plates-back/models"
	"green-plates-back/token"
	"net/http"
)

func GetUserById(c *gin.Context) {
	userId := c.Param("id")

	uniqueId, err := uuid.Parse(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	var user models.User
	result := initializers.DB.First(&user, "id = ?", uniqueId)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
	var users []models.User

	result := initializers.DB.Find(&users)
	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"users": users,
	})
}

func DeleteUser(c *gin.Context) {

	session, err := token.ParseToken(c)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var user models.User
	result := initializers.DB.First(&user, "id = ?", session.Bearer)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	result = initializers.DB.Delete(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.SetCookie("Authorization", "", -1, "", "", false, true)

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Account deleted.",
	})
}
