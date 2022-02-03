package mealitem

import "testing"

func TestNewMealItem(t *testing.T) {
	name := "Tuna Pizza"
	values := []float64{428.6, 20.0, 13.0, 50.0}
	expected := map[float64]string{428.6: "429 kcal", 20.0: "20.0 g", 13.0: "13.0 g", 50.0: "50.0 g"}

	item := NewMealItem(name, values[0], values[1], values[2], values[3])

	if item.Name != name {
		t.Fatalf("name wrong. expected=%q, got=%q", "Tuna Pizza", item.Name)
	}

	if item.Nutrition.Calories != expected[values[0]] {
		t.Fatalf("calories wrong. expected=%q, got=%q", expected[values[0]], item.Nutrition.Calories)
	}
	if item.Nutrition.CarbohydrateContent != expected[values[1]] {
		t.Fatalf("carbohydrateContent wrong. expected=%q, got=%q", expected[values[1]], item.Nutrition.CarbohydrateContent)
	}
	if item.Nutrition.FatContent != expected[values[2]] {
		t.Fatalf("fatContent wrong. expected=%q, got=%q", expected[values[2]], item.Nutrition.FatContent)
	}
	if item.Nutrition.ProteinContent != expected[values[3]] {
		t.Fatalf("proteinContent wrong. expected=%q, got=%q", expected[values[3]], item.Nutrition.ProteinContent)
	}
}
