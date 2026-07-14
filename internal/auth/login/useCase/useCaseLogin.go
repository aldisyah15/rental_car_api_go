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

	u, _ := r.repository.Login(*user)
	return u, nil
}
