package models

import "time"

type Recipe struct {
	ID          uint `gorm:"primaryKey;autoIncrement;"`
	CreatedAt   time.Time
	Title       string
	Description string
	ImagePath   string
	PrepTime    uint
	CookTime    uint
	Servings    uint
	CategoryID  uint
	Steps       string
	Ingredients []RecipeIngredient `gorm:"foreignKey:RecipeID"`
}
