package repository

import (
	"database/sql"
	"rental_car/internal/auth/login/model"

	"golang.org/x/crypto/bcrypt"
)

type LoginRepository struct {
	db *sql.DB
}

func NewLoginRepository(db *sql.DB) *LoginRepository {
	return &LoginRepository{
		db: db,
	}
}

func (r *LoginRepository) Login(req model.RequestLogin) (*model.User, error) {
	var u model.User
	query := "SELECT id, name, username, email, phone, password, rule FROM user WHERE username = ?"
	err := r.db.QueryRow(query, req.Username).Scan(&u.Id, &u.Name, &u.UserName, &u.Email, &u.Phone, &u.Password, &u.Rule)
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password))

	return &u, err
}
