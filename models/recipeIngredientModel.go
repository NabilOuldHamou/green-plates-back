package models

type RecipeIngredient struct {
	RecipeID     uint `json:"-"`
	IngredientID uint
	Quantity     uint
	Measurement  string
	Recipe       Recipe     `gorm:"foreignKey:RecipeID" json:"-"`
	Ingredient   Ingredient `gorm:"foreignKey:IngredientID"`
}
