// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package store

import (
	"time"

	"simple-crud/store/types"
)

type Dosing struct {
	RecipeID     string     `json:"recipe_id"`
	IngredientID string     `json:"ingredient_id"`
	Quantity     int64      `json:"quantity" validate:"required,gt=0"`
	Unit         types.Unit `json:"unit" validate:"required"`
}

type Ingredient struct {
	ID               string         `json:"id"`
	CreatedAt        time.Time      `json:"created_at"`
	Name             string         `json:"name"`
	Description      string         `json:"description"`
	DefaultUnit      string         `json:"default_unit"`
	Category         types.Category `json:"category"`
	AvailableAllYear bool           `json:"available_all_year"`
	AvailableJan     bool           `json:"available_jan"`
	AvailableFeb     bool           `json:"available_feb"`
	AvailableMar     bool           `json:"available_mar"`
	AvailableApr     bool           `json:"available_apr"`
	AvailableMay     bool           `json:"available_may"`
	AvailableJun     bool           `json:"available_jun"`
	AvailableJul     bool           `json:"available_jul"`
	AvailableAug     bool           `json:"available_aug"`
	AvailableSep     bool           `json:"available_sep"`
	AvailableOct     bool           `json:"available_oct"`
	AvailableNov     bool           `json:"available_nov"`
	AvailableDec     bool           `json:"available_dec"`
}

type Recipe struct {
	ID           string    `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Instructions string    `json:"instructions"`
}
