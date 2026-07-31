package model

type RequestFavorite struct {
	Username string `json:"username"`
	IdCar    int    `json:"idCar" binding:"required"`
}

type ResponseFavorite struct {
	Username string `json:"username"`
	IdCar    int    `json:"idCar"`
	CreateAt string `json:"createAt"`
}
