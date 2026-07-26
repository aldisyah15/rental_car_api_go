package repository

import (
	"database/sql"
	"rental_car/internal/auth/register/model"

	"golang.org/x/crypto/bcrypt"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(User *model.User) error {
	query := "INSERT INTO user (name, username, email, password, phone) VALUES (?, ?, ?, ?, ?)"
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(User.Password), 14)
	_, err := r.db.Exec(query, User.Name, User.UserName, User.Email, hashPassword, User.Phone)
	if err != nil {
		return err
	}
	return err
}

func (r *AuthRepository) CheckUserAndEmailExists(username string, email string) (bool, bool, error) {
	var isExistUsername bool
	var isExistEmail bool

	query := "select exists(select 1 from user where username = ?)"
	err := r.db.QueryRow(query, username).Scan(&isExistUsername)
	if err != nil {
		return false, false, err
	}

	query = "select exists(select 1 from user where email = ?)"
	err = r.db.QueryRow(query, email).Scan(&isExistEmail)
	if err != nil {
		return false, false, err
	}

	return isExistUsername, isExistEmail, err
}
