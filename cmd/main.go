package main

import (
	"log"
	"math/rand"

	"github.com/cukiprit/anekapay-test-case/internal/animal"
	"github.com/cukiprit/anekapay-test-case/internal/db"
	"github.com/gin-gonic/gin"
)

func SeedAnimals() {
	animalNames := []string{"Lion", "Tiger", "Bear", "Elephant", "Giraffe", "Zebra", "Kangaroo", "Panda", "Wolf", "Fox"}
	animalClasses := []string{"Mammal", "Bird", "Reptile", "Amphibian", "Fish"}

	for i := 1; i <= 50; i++ {
		newAnimal := animal.Animal{
			ID:    i,
			Name:  animalNames[rand.Intn(len(animalNames))],
			Class: animalClasses[rand.Intn(len(animalClasses))],
			Legs:  rand.Intn(5) + 2,
		}

		_, err := db.Database.Exec("INSERT INTO animals (id, name, class, legs) VALUES (?, ?, ?, ?)", newAnimal.ID, newAnimal.Name, newAnimal.Class, newAnimal.Legs)
		if err != nil {
			log.Printf("Failed to insert animal %d: %s", i, err)
		}
	}
}

func main() {
	db.InitDB()
	defer db.CloseDB()

	// Comment this line if it's already execute once
	// SeedAnimals()

	r := gin.Default()

	animal.SetupRoutes(r)

	r.Run()
}
