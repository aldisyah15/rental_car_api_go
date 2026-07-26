package useCase

import (
	"errors"
	"rental_car/internal/auth/login/model"
	"rental_car/internal/auth/login/repository"
)

type Login struct {
	repository *repository.LoginRepository
}

func NewUseCaseLogin(r *repository.LoginRepository) *Login {
	return &Login{
		repository: r,
	}
}

func (r *Login) AuthLogin(req *model.RequestLogin) (*model.User, error) {
	user := &model.RequestLogin{
		Username: req.Username,
		Password: req.Password,
	}

	result, err := r.repository.Login(*user)

	if err != nil {
		return nil, errors.New("username or password wrong")
	}

	return result, err
}
