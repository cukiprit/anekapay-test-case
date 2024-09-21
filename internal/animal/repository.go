package animal

import (
	"database/sql"
	"errors"
)

type DBRepository struct {
	db *sql.DB
}

func NewDBRepository(db *sql.DB) *DBRepository {
	return &DBRepository{db: db}
}

func (r *DBRepository) Create(animal *Animal) error {
	var exists int
	err := r.db.QueryRow("SELECT COUNT(*) FROM animals WHERE id = ?", animal.ID).Scan(&exists)
	if err != nil {
		return err
	}
	if exists > 0 {
		return errors.New("Animal with this ID already exists")
	}

	_, err = r.db.Exec("INSERT INTO animals (id, name, class, legs) VALUES (?, ?, ?, ?)", animal.ID, animal.Name, animal.Class, animal.Legs)

	return err
}

func (r *DBRepository) Update(animal *Animal) error {
	_, err := r.db.Exec("UPDATE animals SET name = ?, class = ?, legs = ? WHERE id = ?", animal.Name, animal.Class, animal.Legs, animal.ID)
	return err
}

func (r *DBRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM animals WHERE id = ?", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("animal not found")
	}
	return nil
}

func (r *DBRepository) GetAll() ([]Animal, error) {
	rows, err := r.db.Query("SELECT id, name, class, legs FROM animals")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var animals []Animal
	for rows.Next() {
		var animal Animal
		if err := rows.Scan(&animal.ID, &animal.Name, &animal.Class, &animal.Legs); err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}
	return animals, nil
}

func (r *DBRepository) GetByID(id int) (*Animal, error) {
	var animal Animal
	err := r.db.QueryRow("SELECT id, name, class, legs FROM animals WHERE id = ?", id).Scan(&animal.ID, &animal.Name, &animal.Class, &animal.Legs)
	if err == sql.ErrNoRows {
		return nil, errors.New("animal not found")
	}
	return &animal, err
}
