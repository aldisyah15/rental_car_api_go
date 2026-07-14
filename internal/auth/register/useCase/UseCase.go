package useCase

import (
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

func (r *AuthCase) CreateUser(req model.RegisterRequest) error {
	user := &model.User{
		UserName: req.UserName,
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
	}

	return r.repository.CreateUser(user)
}
