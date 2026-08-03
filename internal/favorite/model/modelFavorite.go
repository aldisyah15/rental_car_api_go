package model

type RequestFavorite struct {
	Username string `json:"username"`
	IdCar    int    `json:"idCar" binding:"required"`
}

type ResponseFavorite struct {
	Username    string   `json:"username"`
	IdCar       int      `json:"idCar"`
	CreateAt    string   `json:"createAt"`
	Name        string   `json:"name"`
	Brand       string   `json:"brand"`
	RentalPrice float64  `json:"rental_price"`
	Images      []string `json:"images"`
	Horsepower  string   `json:"horse_power"`
	Gear        string   `json:"gear"`
	Description string   `json:"description"`
	Seat        string   `json:"seat"`
	Stock       int      `json:"stock"`
	Logo        string   `json:"logo"`
}
