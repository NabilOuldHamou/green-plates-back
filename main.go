package main

import (
	"github.com/gin-gonic/gin"
	"green-plates-back/api"
	"green-plates-back/controllers"
	"green-plates-back/initializers"
	"green-plates-back/middleware"
	"log"
	"os"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	api.CreateRouter()

	// Auth
	api.V1.POST("/auth/signup", controllers.Signup)
	api.V1.POST("/auth/login", controllers.Login)
	api.V1.POST("/auth/logout", controllers.Logout)

	// Users
	api.V1.GET("/users", controllers.GetUsers)
	api.V1.GET("/users/:id", controllers.GetUserById)
	api.V1.DELETE("/users", middleware.RequireAuth, controllers.DeleteUser)

	// Ingredients
	api.V1.GET("/ingredients", controllers.GetIngredients)

	// Recipes
	api.V1.GET("/recipes", controllers.GetRecipes)
	api.V1.POST("/recipes", middleware.RequireAuth, controllers.CreateRecipe)

	err := api.Router.Run()
	if err != nil {
		log.Fatal("Router could not be created.")
	}
}
