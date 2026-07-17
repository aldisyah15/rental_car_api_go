package useCase

import (
	"log"
	"rental_car/internal/user/model"
	"rental_car/internal/user/repository"
)

type UserUseCase struct {
	repository *repository.UserRepository
}

func NewUserUseCase(username *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		repository: username,
	}
}

func (u UserUseCase) GetUser(username string) (*model.RespondUser, error) {
	user := &model.RequestUser{
		Username: username,
	}
	result, err := u.repository.GetUser(user.Username)
	return result, err
}

func (u UserUseCase) UpdateUser(username string, req model.RequestUpdateUser) (*model.RespondUser, error) {
	existingUser, err := u.repository.GetUser(username)
	if err != nil {
		return nil, err
	}

	log.Printf("log user useCase: %v", existingUser)

	if req.Name != "" {
		existingUser.Name = req.Name
	}
	if req.Username != "" {
		existingUser.UserName = req.Username
	}
	if req.Phone != "" {
		existingUser.Phone = req.Phone
	}
	if req.Email != "" {
		existingUser.Email = req.Email
	}

	user := &model.RequestUpdateUser{
		Name:     existingUser.Name,
		Username: existingUser.UserName,
		Phone:    existingUser.Phone,
		Email:    existingUser.Email,
	}
	return u.repository.UpdateUser(username, user)
}

func (u UserUseCase) DeleteUser(username string) (id string, err error) {
	id, err = u.repository.DeleteUser(username)
	return id, err
}
