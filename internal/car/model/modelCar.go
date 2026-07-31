package model

type RequestCar struct {
	Name        string  `form:"name"`
	Brand       string  `form:"brand"`
	RentalPrice float64 `form:"rental_price"`
	HorsePower  string  `form:"horse_power"`
	Gear        string  `form:"gear"`
	Description string  `form:"description"`
	Seat        string  `form:"seat"`
	Stock       int     `form:"stock"`
}

type ResponseCar struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Brand       string   `json:"brand"`
	RentalPrice float64  `json:"rental_price"`
	Images      []string `json:"images"`
	Horsepower  string   `json:"horse_power"`
	Gear        string   `json:"gear"`
	Description string   `json:"description"`
	Seat        string   `json:"seat"`
	Stock       int      `json:"stock"`
	Logo        string   `json:"brand"`
}
