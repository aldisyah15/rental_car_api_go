package model

type RespondUser struct {
	ID       int    `json:"ID"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"username"`
	Phone    string `json:"phone"`
}

type RequestUser struct {
	Username string `json:"userName"`
}
