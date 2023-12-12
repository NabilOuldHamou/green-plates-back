package models

import "time"

type Recipe struct {
	ID          uint `gorm:"primaryKey;autoIncrement;"`
	CreatedAt   time.Time
	Title       string
	Description string
	PrepTime    uint
	CookTime    uint
	Servings    uint
	CategoryID  uint
	Steps       []Step
	Ingredients []RecipeIngredient `gorm:"foreignKey:RecipeID`
}
