/*
Package food implements classes that are used for meal planning and calorie
counting.

It will reuse the naming conventions established by schema.org.

- https://schema.org/Recipe
- https://schema.org/NutritionInformation
*/

package food

import (
	"fmt"
	"math"
)

// https://schema.org/NutritionInformation
type MealItem struct {
	name      string
	nutrition *Nutrition
}

type Nutrition struct {
	calories            string // <Number> <Energy unit of measure>
	carbohydrateContent string // <Number> <Mass unit of measure> -> in gram
	fatContent          string // <Number> <Mass unit of measure> -> in gram
	proteinContent      string // <Number> <Mass unit of measure> -> in gram
}

func NewMealItem(name string, calories, carbs, fat, protein float64) *MealItem {
	item := &MealItem{
		name: name,
		nutrition: &Nutrition{
			calories:            fmt.Sprintf("%d kcal", int(math.Round(calories))),
			fatContent:          fmt.Sprintf("%.1f g", fat),
			carbohydrateContent: fmt.Sprintf("%.1f g", carbs),
			proteinContent:      fmt.Sprintf("%.1f g", protein),
		},
	}
	return item
}
