package controllers

import (
	"github.com/gin-gonic/gin"
	"green-plates-back/initializers"
	"green-plates-back/models"
	"net/http"
)

func GetRecipes(c *gin.Context) {
	var recipes []models.Recipe

	result := initializers.DB.Model(&models.Recipe{}).Preload("Steps").Preload("Ingredients").Preload("Ingredients.Ingredient").Find(&recipes)
	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"recipes": recipes,
	})
}

func CreateRecipe(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	result := initializers.DB.Create(&recipe)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "An internal server error occurred",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"recipeId": recipe.ID,
	})
}
