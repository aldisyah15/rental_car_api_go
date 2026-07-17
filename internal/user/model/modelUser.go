package model

type RespondUser struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"username"`
	Phone    string `json:"phone"`
}

type RequestUser struct {
	Username string `json:"userName"`
}

type RequestUpdateUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
