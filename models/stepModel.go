package models

type Step struct {
	ID          uint `json:"-"`
	Description string
	RecipeID    uint `json:"-"`
}
