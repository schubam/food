package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/schubam/food/meal_item"
	"log"
	"os"
)

const fileName = "sqlite.db"

func main() {
	os.Remove(fileName)

	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
	}

	mealItemsRepository := mealitem.NewSQLiteRepository(db)
	if err := mealItemsRepository.Migrate(); err != nil {
		log.Fatalf("error migrating db: %v\n", err)
	}

	tunasalad := mealitem.MealItem{
		Name:     "Tuna Salad",
		Calories: 420,
		Carbs:    123,
		Fat:      12,
		Protein:  69,
	}

	broccolistew := mealitem.MealItem{
		Name:     "Broccoli Casserole",
		Calories: 469,
		Carbs:    312,
		Fat:      21,
		Protein:  96,
	}

	createdTunasalad, err := mealItemsRepository.Create(tunasalad)
	if err != nil {
		log.Fatalf("error Create(): %v\n", err)
	}

	createdBroccoli, err := mealItemsRepository.Create(broccolistew)
	if err != nil {
		log.Fatalf("error Create(): %v\n", err)
	}

	gotTunasalad, err := mealItemsRepository.GetByName("Tuna Salad")
	if err != nil {
		log.Fatalf("error GetByName(): %v\n", err)
	}

	fmt.Printf("get by name: %+v\n", gotTunasalad)

	createdTunasalad.Calories = 666
	if _, err := mealItemsRepository.Update(createdTunasalad.ID, *createdTunasalad); err != nil {
		log.Fatalf("error Update(): %v\n", err)
	}

	all, err := mealItemsRepository.All()
	if err != nil {
		log.Fatalf("error All(): %v\n", err)
	}

	fmt.Printf("\nAll mealitems:\n")
	for _, mealitem := range all {
		fmt.Printf("mealitem: %+v\n", mealitem)
	}

	if err := mealItemsRepository.Delete(createdBroccoli.ID); err != nil {
		log.Fatalf("error Delete(): %v\n", err)
	}

	all, err = mealItemsRepository.All()
	if err != nil {
		log.Fatalf("error All(): %v\n", err)
	}
	fmt.Printf("\nAll mealitems:\n")
	for _, mealitem := range all {
		fmt.Printf("mealitem: %+v\n", mealitem)
	}
}
