package repository

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"rental_car/internal/booking/model"
)

type BookingRepository struct {
	db   *sql.DB
	http *http.Client
}

func NewRepositoryBooking(db *sql.DB, http *http.Client) *BookingRepository {
	return &BookingRepository{db: db, http: http}
}

func (r *BookingRepository) ProsesToCheckout(req *model.Booking) (*model.Booking, error) {
	queryInsert := `INSERT INTO bookings
    (username, id_car, pickup_date,drop_date, days, discount, tax, status, total_rental_price, grand_total)
	values (?,?,?,?,?,?,?,?,?,?)`

	rows, err := r.db.Exec(
		queryInsert,
		req.Username,
		req.IdCar,
		req.PickupDate,
		req.DropDate,
		req.Days,
		req.Discount,
		req.Tax,
		req.Status,
		req.TotalRentalPrice,
		req.GrandTotal)
	if err != nil {
		return nil, err
	}

	idOrder, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	var resultBooking model.Booking
	querySelect := `SELECT id_order,username, id_car, pickup_date,drop_date, days, discount, tax, status, createAt, total_rental_price, grand_total from bookings where id_order = ? and username = ?`
	err = r.db.QueryRow(querySelect, idOrder, req.Username).Scan(
		&resultBooking.IdOrder,
		&resultBooking.Username,
		&resultBooking.IdCar,
		&resultBooking.PickupDate,
		&resultBooking.DropDate,
		&resultBooking.Days,
		&resultBooking.Discount,
		&resultBooking.Tax,
		&resultBooking.Status,
		&resultBooking.CreateAt,
		&resultBooking.TotalRentalPrice,
		&resultBooking.GrandTotal,
	)
	if err != nil {
		return nil, err
	}
	return &resultBooking, nil
}

func (r *BookingRepository) GetDetailCheckout(idOrder int, username string) (*model.Booking, error) {
	querySelect := `SELECT id_order,username, id_car, pickup_date,drop_date, days, discount, tax, status, createAt, total_rental_price, grand_total from bookings where id_order = ? and username = ?`
	var resultBooking model.Booking
	err := r.db.QueryRow(querySelect, idOrder, username).Scan(
		&resultBooking.IdOrder,
		&resultBooking.Username,
		&resultBooking.IdCar,
		&resultBooking.PickupDate,
		&resultBooking.DropDate,
		&resultBooking.Days,
		&resultBooking.Discount,
		&resultBooking.Tax,
		&resultBooking.Status,
		&resultBooking.CreateAt,
		&resultBooking.TotalRentalPrice,
		&resultBooking.GrandTotal,
	)

	if err != nil {
		return nil, err
	}

	return &resultBooking, nil
}

func (r *BookingRepository) BookingNow(req *model.PaymentRequest) (*model.GopayChargeResponse, error) {
	url := "https://api.sandbox.midtrans.com/v2/charge"

	payload := &model.PaymentRequest{
		PaymentType: req.PaymentType,
		TransactionDetails: model.TransactionDetails{
			OrderId:     req.TransactionDetails.OrderId,
			GrossAmount: req.TransactionDetails.GrossAmount,
		},
	}

	jsonByte, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonByte))
	if err != nil {
		return nil, err
	}

	authString := base64.StdEncoding.EncodeToString([]byte("SB-Mid-server-9lx-ld7pM62I44yhYCgdSWKe" + ":"))
	httpReq.Header.Set("Authorization", "Basic "+authString)
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Basic "+authString)

	resp, err := r.http.Do(httpReq)
	if err != nil {
		return nil, err
	}

	log.Printf("resp: %v", resp.StatusCode)
	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	log.Printf("resp: %v", string(bodyByte))

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("midtrans error: %s", string(bodyByte))
	}
	var bookingResp model.GopayChargeResponse
	err = json.Unmarshal(bodyByte, &bookingResp)
	if err != nil {
		return nil, err
	}

	return &bookingResp, nil
}
