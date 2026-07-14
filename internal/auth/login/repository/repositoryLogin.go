package repository

import (
	"database/sql"
	"errors"
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
	query := "SELECT id, name, username, email, password from TABLE ( user ) where user = ?"

	err := r.db.QueryRow(query, req.Username).Scan(&u.Name, &u.UserName, &u.Email, &u.Phone, &u.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Password or Email is wrong!")
		}
		return nil, nil
	}

	hastPassword := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password))
	if hastPassword != nil {
		return nil, errors.New("email atau password salah")
	}

	return &u, nil
}
