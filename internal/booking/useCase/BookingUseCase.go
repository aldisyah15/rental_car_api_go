package useCase

import (
	"database/sql"
	"errors"
	"log"
	"rental_car/internal/booking/model"
	"rental_car/internal/booking/repository"
	repository2 "rental_car/internal/car/repository"
	"strconv"
	"strings"
	"time"
)

type BookingUseCase struct {
	db    *repository.BookingRepository
	dbCar *repository2.CarRepository
}

func NewUseCaseBooking(repo *repository.BookingRepository, repoCar *repository2.CarRepository) *BookingUseCase {
	return &BookingUseCase{db: repo, dbCar: repoCar}
}

func (c BookingUseCase) ProsesToCheckout(req *model.RequestProsesToCheckout) (*model.Booking, error) {
	const taxPersen = 12.0
	const discount = 50.0

	pickupDate, err := time.Parse("2006-01-02", req.PickupDate)
	dropDate, err := time.Parse("2006-01-02", req.DropDate)
	if err != nil {
		return nil, err
	}

	pickupDateFormatted := pickupDate.Format("2006-01-02")
	dropDateFormatted := dropDate.Format("2006-01-02")
	days := int(dropDate.Sub(pickupDate).Hours() / 24)

	req = &model.RequestProsesToCheckout{
		PickupDate: pickupDateFormatted,
		DropDate:   dropDateFormatted,
		IdCar:      req.IdCar,
		Username:   req.Username,
	}
	repoCar, err := c.dbCar.GetCarById(req.IdCar)
	if err != nil {
		return nil, err
	}

	totalRentalPrice := float64(days) * repoCar.RentalPrice
	taxAmount := (totalRentalPrice * taxPersen) / 100
	discountAmount := (totalRentalPrice * discount) / 100
	grandTotal := totalRentalPrice + taxAmount - discountAmount

	result := &model.Booking{
		PickupDate:       req.PickupDate,
		DropDate:         req.DropDate,
		IdCar:            req.IdCar,
		Username:         req.Username,
		Days:             days,
		Discount:         discount,
		Tax:              taxPersen,
		Status:           "draft",
		TotalRentalPrice: totalRentalPrice,
		GrandTotal:       grandTotal,
	}

	repo, err := c.db.ProsesToCheckout(result)
	formattedDate := strings.Split(repo.PickupDate, "T")[0]

	log.Printf("useCase date in booking : %s", formattedDate)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (c BookingUseCase) GetDetailCheckout(idOrder int, username string) (*model.Booking, error) {
	result, err := c.db.GetDetailCheckout(idOrder, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("booking not found")
		}
	}

	return result, nil
}

func (c BookingUseCase) BookingNow(idOrder string, username string) (*model.GopayChargeResponse, error) {
	idInt, _ := strconv.Atoi(idOrder)
	_, err := c.db.GetDetailCheckout(idInt, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("booking not found")
		}
	}
	req := &model.PaymentRequest{
		PaymentType: "gopay",
		TransactionDetails: model.TransactionDetails{
			OrderId:     idOrder,
			GrossAmount: 100000,
		},
	}
	res, err := c.db.BookingNow(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
