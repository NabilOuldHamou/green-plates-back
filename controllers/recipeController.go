package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"green-plates-back/initializers"
	"green-plates-back/models"
	"net/http"
	"strings"
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
	var body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		PrepTime    uint   `json:"preptime"`
		CookTime    uint   `json:"cooktime"`
		Servings    uint   `json:"servings"`
		CategoryID  uint   `json:"category_id"`
		Steps       string `json:"steps"`
	}

	file, _ := c.FormFile("file")
	splitName := strings.Split(file.Filename, ".")
	path := "assets/recipes/" + uuid.New().String() + "." + splitName[len(splitName)-1]
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown error"})
		return
	}

	// TODO resize file & change format to JPEG

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ingredientsJson := c.PostForm("Ingredients")
	var ingredients []models.RecipeIngredient
	if err := json.Unmarshal([]byte(ingredientsJson), &ingredients); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ingredients format"})
		return
	}

	recipe := models.Recipe{
		Title:       body.Title,
		Description: body.Description,
		ImagePath:   path,
		PrepTime:    body.PrepTime,
		CookTime:    body.CookTime,
		Servings:    body.Servings,
		CategoryID:  body.CategoryID,
		Steps:       body.Steps,
		Ingredients: ingredients,
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
