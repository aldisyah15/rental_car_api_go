package repository

import (
	"database/sql"
	"errors"
	"log"
	"rental_car/internal/user/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) GetUser(userName string) (*model.RespondUser, error) {
	var u model.RespondUser
	query := "SELECT id, name, username, email, phone from user where username = ?"
	err := r.db.QueryRow(query, userName).Scan(&u.ID, &u.Name, &u.UserName, &u.Email, &u.Phone)

	log.Printf("username: %v", userName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &u, nil
}
