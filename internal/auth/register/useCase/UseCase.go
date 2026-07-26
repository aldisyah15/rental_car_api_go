package useCase

import (
	"errors"
	"rental_car/internal/auth/register/model"
	"rental_car/internal/auth/register/repository"
)

type AuthCase struct {
	repository *repository.AuthRepository
}

func NewAuthCase(repo *repository.AuthRepository) *AuthCase {
	return &AuthCase{
		repository: repo,
	}
}

func (r *AuthCase) CreateUser(req *model.RegisterRequest) error {
	username, email, err := r.repository.CheckUserAndEmailExists(req.UserName, req.Email)

	if username == true {
		return errors.New("username already exists")
	}
	if email == true {
		return errors.New("email already exists")
	}

	if err != nil {
		return err
	}
	user := &model.User{
		Name:     req.UserName,
		UserName: req.UserName,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
	}

	result := r.repository.CreateUser(user)
	return result
}
