package model

type RequestFavorite struct {
	Username string `json:"username"`
	IdCar    int    `json:"idCar" binding:"required"`
}
