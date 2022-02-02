package food

import "testing"

func TestNewMealItem(t *testing.T) {
	name := "Tuna Pizza"
	values := []float64{428.6, 20.0, 13.0, 50.0}
	expected := map[float64]string{428.6: "429 kcal", 20.0: "20.0 g", 13.0: "13.0 g", 50.0: "50.0 g"}

	item := NewMealItem(name, values[0], values[1], values[2], values[3])

	if item.name != name {
		t.Fatalf("name wrong. expected=%q, got=%q", "Tuna Pizza", item.name)
	}

	if item.nutrition.calories != expected[values[0]] {
		t.Fatalf("calories wrong. expected=%q, got=%q", expected[values[0]], item.nutrition.calories)
	}
	if item.nutrition.carbohydrateContent != expected[values[1]] {
		t.Fatalf("carbohydrateContent wrong. expected=%q, got=%q", expected[values[1]], item.nutrition.carbohydrateContent)
	}
	if item.nutrition.fatContent != expected[values[2]] {
		t.Fatalf("fatContent wrong. expected=%q, got=%q", expected[values[2]], item.nutrition.fatContent)
	}
	if item.nutrition.proteinContent != expected[values[3]] {
		t.Fatalf("proteinContent wrong. expected=%q, got=%q", expected[values[3]], item.nutrition.proteinContent)
	}
}
