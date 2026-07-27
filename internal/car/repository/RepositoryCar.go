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

func (r CarRepository) GetAllCars() (*[]model.ResponseCar, error) {

	query := "SELECT id, name, brand, rental_price, image, horse_power, gear, description, stock from detail_car"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []model.ResponseCar
	for rows.Next() {
		var rawJSON []byte
		var car model.ResponseCar

		_ = rows.Scan(
			&car.Id,
			&car.Name,
			&car.Brand,
			&car.RentalPrice,
			&rawJSON,
			&car.Horsepower,
			&car.Gear,
			&car.Description,
			&car.Stock,
		)

		if len(rawJSON) > 0 {
			_ = json.Unmarshal(rawJSON, &car.Images)
		}
		cars = append(cars, car)
	}

	return &cars, nil
}
