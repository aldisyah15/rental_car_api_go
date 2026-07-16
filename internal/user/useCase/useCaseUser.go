package useCase

import (
	"rental_car/internal/user/model"
	"rental_car/internal/user/repository"
)

type UserUseCase struct {
	repository *repository.UserRepository
}

func NewUserUseCase(id *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		repository: id,
	}
}

func (u UserUseCase) GetUser(username string) (*model.RespondUser, error) {
	user := &model.RequestUser{
		Username: username,
	}
	result, err := u.repository.GetUser(user.Username)
	return result, err
}
