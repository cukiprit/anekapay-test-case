package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/cukiprit/anekapay-test-case/db"
	"github.com/cukiprit/anekapay-test-case/models"
	"github.com/gin-gonic/gin"
)

func CreateAnimal(ctx *gin.Context) {
	var newAnimal models.Animal

	if err := ctx.ShouldBindJSON(&newAnimal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input!"})
		return
	}

	var exists int
	err := db.Database.QueryRow("SELECT COUNT(*) FROM animals WHERE id = ?", newAnimal.ID).Scan(&exists)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if exists > 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Animal with this ID already exists."})
	}

	_, err = db.Database.Exec("INSERT INTO animals (id, name, class, legs) VALUES (?, ?, ?, ?)", newAnimal.ID, newAnimal.Name, newAnimal.Class, newAnimal.Legs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
	}

	ctx.JSON(http.StatusCreated, newAnimal)
}

func UpdateAnimal(ctx *gin.Context) {
	var updatedAnimal models.Animal
	if err := ctx.ShouldBindJSON(&updatedAnimal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := db.Database.Exec("UPDATE animals SET name = ?, class = ?, legs = ?, WHERE id = ?", updatedAnimal.Name, updatedAnimal.Class, updatedAnimal.Legs, updatedAnimal.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err := db.Database.Exec("INSERT INTO animals (id, name, class, legs) VALUES (?, ?, ?, ?)", updatedAnimal.ID, updatedAnimal.Name, updatedAnimal.Class, updatedAnimal.Legs)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
				return
			}
			ctx.JSON(http.StatusCreated, updatedAnimal)
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	ctx.JSON(http.StatusOK, updatedAnimal)
}

func DeleteAnimal(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	_, err = db.Database.Exec("DELETE FROM animals WHERE id = ?", id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Animal not found"})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func GetAllAnimals(ctx *gin.Context) {
	rows, err := db.Database.Query("SELECT id, name, class, legs FROM animals")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	var animals []models.Animal
	for rows.Next() {
		var animal models.Animal
		if err := rows.Scan(&animal.ID, &animal.Name, &animal.Class, &animal.Legs); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}
		animals = append(animals, animal)
	}

	if len(animals) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No animals found."})
		return
	}
	ctx.JSON(http.StatusOK, animals)
}

func GetAnimalByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var animal models.Animal
	err = db.Database.QueryRow("SELECT id, name, class, legs FROM animals WHERE id = ?", id).Scan(&animal.ID, &animal.Name, &animal.Class, &animal.Legs)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Animal not found."})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	ctx.JSON(http.StatusOK, animal)
}
