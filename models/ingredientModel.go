package models

type Ingredient struct {
	ID      uint               `gorm:"primaryKey;autoIncrement;"`
	Name    string             `gorm:"unique"`
	Recipes []RecipeIngredient `gorm:"foreignKey:IngredientID" json:"-"`
}
