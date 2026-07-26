package useCase

import (
	"rental_car/internal/car/model"
	"rental_car/internal/car/repository"
)

type CarUseCase struct {
	repository *repository.CarRepository
}

func NewCarUseCase(car *repository.CarRepository) *CarUseCase {
	return &CarUseCase{car}
}

func (c CarUseCase) UploadRentalCar(car *model.RequestCar, urlImages []string) error {
	cars := &model.RequestCar{
		Name:        car.Name,
		Brand:       car.Brand,
		RentalPrice: car.RentalPrice,
		HorsePower:  car.HorsePower,
		Description: car.Description,
		Gear:        car.Gear,
		Seat:        car.Seat,
		Stock:       car.Stock,
	}

	err := c.repository.UploadRentalCar(cars, urlImages)

	return err
}
