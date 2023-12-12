package controllers

import (
	"github.com/gin-gonic/gin"
	"green-plates-back/initializers"
	"green-plates-back/models"
	"net/http"
)

func GetIngredients(c *gin.Context) {
	var ingredients []models.Ingredient

	result := initializers.DB.Find(&ingredients)
	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"ingredients": ingredients,
	})
}
