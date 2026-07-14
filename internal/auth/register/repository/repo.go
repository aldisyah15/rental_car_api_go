package repository

import (
	"database/sql"
	"log"
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
	result, err := r.db.Exec(query, User.Name, User.UserName, User.Email, hashPassword, User.Phone)

	if err != nil {
		log.Printf("Gagal masuk ke database", err)
	}
	id, _ := result.LastInsertId()
	if result != nil {
		log.Printf("id %v berhasil masuk ke database", id)
	}
	return nil
}
