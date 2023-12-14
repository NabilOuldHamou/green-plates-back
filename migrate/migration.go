package main

import (
	"green-plates-back/initializers"
	"green-plates-back/models"
	"log"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	log.Println("Migrating models to db...")

	err := initializers.DB.AutoMigrate(
		&models.User{},
		&models.Ingredient{},
		&models.Category{},
		&models.Recipe{},
		&models.RecipeIngredient{})
	if err != nil {
		log.Fatalf("Automatic migration has failed : %v", err)
	}

	log.Println("Migration successful.")
}
