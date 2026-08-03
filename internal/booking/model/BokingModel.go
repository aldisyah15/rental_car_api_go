package model

type RequestProsesToCheckout struct {
	PickupDate string `json:"pickupDate" binding:"required"`
	DropDate   string `json:"dropDate" binding:"required"`
	IdCar      int    `json:"id_Car" binding:"required"`
	Username   string `json:"username"`
}

type Booking struct {
	IdOrder          int     `json:"id_Order" binding:"required"`
	PickupDate       string  `json:"pickupDate" binding:"required"`
	DropDate         string  `json:"dropDate" binding:"required"`
	IdCar            int     `json:"id_Car" binding:"required"`
	Username         string  `json:"username" binding:"required"`
	Days             int     `json:"days" binding:"required"`
	Discount         int     `json:"discount" binding:"required"`
	Tax              int     `json:"tax" binding:"required"`
	Status           string  `json:"status" binding:"required"`
	CreateAt         string  `json:"createAt" binding:"required"`
	TotalRentalPrice float64 `json:"total_rental_price" binding:"required"`
	GrandTotal       float64 `json:"grandTotal" binding:"required"`
}

type PaymentRequest struct {
	PaymentType        string             `json:"payment_type"`
	TransactionDetails TransactionDetails `json:"transaction_details"`
}

type TransactionDetails struct {
	OrderId     string `json:"order_id"`
	GrossAmount int64  `json:"gross_amount"`
}

type Action struct {
	Name   string `json:"name"`
	Method string `json:"method"`
	URL    string `json:"url"`
}

type GopayChargeResponse struct {
	StatusCode        string   `json:"status_code"`
	StatusMessage     string   `json:"status_message"`
	TransactionID     string   `json:"transaction_id"`
	OrderID           string   `json:"order_id"`
	MerchantID        string   `json:"merchant_id"`
	GrossAmount       string   `json:"gross_amount"`
	Currency          string   `json:"currency"`
	PaymentType       string   `json:"payment_type"`
	TransactionTime   string   `json:"transaction_time"`
	TransactionStatus string   `json:"transaction_status"`
	FraudStatus       string   `json:"fraud_status"`
	Actions           []Action `json:"actions"`
	QrString          string   `json:"qr_string"`
	ExpiryTime        string   `json:"expiry_time"`
}
