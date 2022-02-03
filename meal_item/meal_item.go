/*
Package food implements classes that are used for meal planning and calorie
counting.

It will reuse the naming conventions established by schema.org.

- https://schema.org/Recipe
- https://schema.org/NutritionInformation
*/

package mealitem

import (
	"fmt"
	"math"
)

// https://schema.org/NutritionInformation
type Nutrition struct {
	Calories            string // <Number> <Energy unit of measure>
	CarbohydrateContent string // <Number> <Mass unit of measure> -> in gram
	FatContent          string // <Number> <Mass unit of measure> -> in gram
	ProteinContent      string // <Number> <Mass unit of measure> -> in gram
}

// https://schema.org/Recipe
type MealItem struct {
	ID        int64
	Name      string
	Calories  int
	Carbs     int
	Fat       int
	Protein   int
	Nutrition *Nutrition
}

func NewMealItem(name string, calories, carbs, fat, protein float64) *MealItem {
	item := &MealItem{
		Name: name,
		Nutrition: &Nutrition{
			Calories:            fmt.Sprintf("%d kcal", int(math.Round(calories))),
			FatContent:          fmt.Sprintf("%.1f g", fat),
			CarbohydrateContent: fmt.Sprintf("%.1f g", carbs),
			ProteinContent:      fmt.Sprintf("%.1f g", protein),
		},
	}
	return item
}
