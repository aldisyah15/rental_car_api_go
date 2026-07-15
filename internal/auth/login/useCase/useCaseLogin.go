package useCase

import (
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

	u, err := r.repository.Login(*user)
	if err != nil {
		return nil, err
	}
	return u, err
}
