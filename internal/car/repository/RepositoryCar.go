package repository

import (
	"database/sql"
	"encoding/json"
	"rental_car/internal/car/model"
)

type CarRepository struct {
	db *sql.DB
}

func NewRepositoryCar(db *sql.DB) *CarRepository {
	return &CarRepository{db: db}
}

func (r CarRepository) UploadRentalCar(car *model.RequestCar, urlImages []string) error {
	urlJson, _ := json.Marshal(urlImages)
	query := "INSERT INTO detail_car (name, brand, rental_price, image, horse_power, gear, description, seat, stock) VALUES (?,?,?, ?,?,?,?,?,?)"
	_, err := r.db.Exec(query, car.Name, car.Brand, car.RentalPrice, string(urlJson), car.HorsePower, car.Gear, car.Description, car.Seat, car.Stock)

	return err
}

func (r CarRepository) GetDetailCar() (*model.ResponseCar, error) {
	var cars model.ResponseCar
	query := "SELECT id, name, brand, rental_price, images, horse_power, gear, description, stock from detail_car"
	err := r.db.QueryRow(query).Scan(&cars.Id, &cars.Name, &cars.Brand, &cars.RentalPrice, cars.Images, cars.Horsepower, &cars.Gear, &cars.Description, &cars.Stock)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
	}

	return &cars, nil
}
