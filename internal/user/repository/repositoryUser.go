package repository

import (
	"database/sql"
	"errors"
	"rental_car/internal/user/model"
	"strconv"
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

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &u, nil
}

func (r UserRepository) UpdateUser(username string, req *model.RequestUpdateUser) (*model.RespondUser, error) {
	var u model.RespondUser
	querySelect := "SELECT id, name, username, email, phone from user where username = ?"
	err := r.db.QueryRow(querySelect, username).Scan(&u.ID, &u.Name, &u.UserName, &u.Email, &u.Phone)
	if err != nil {
		return nil, err
	}

	queryUpdate := "UPDATE user set name = ?, username = ?, email = ?, phone = ? where id = ?"
	_, err = r.db.Exec(queryUpdate, req.Name, req.Username, req.Email, req.Phone, u.ID)
	if err != nil {
		return nil, err
	}

	return &model.RespondUser{
		ID:       u.ID,
		Name:     req.Name,
		UserName: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
	}, nil
}

func (r UserRepository) DeleteUser(username string) (string, error) {
	var id int
	querySelect := "SELECT id from user where username = ?"
	_ = r.db.QueryRow(querySelect, username).Scan(&id)

	queryDelete := "DELETE from user where username = ?"
	_, err := r.db.Exec(queryDelete, username)
	if err != nil {
		return "", errors.New("user not found")
	}

	return strconv.Itoa(id), err
}
