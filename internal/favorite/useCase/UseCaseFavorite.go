package useCase

import (
	"rental_car/internal/favorite/model"
	"rental_car/internal/favorite/repository"
)

type FavoriteUseCase struct {
	repository *repository.FavoriteRepository
}

func NewUseCaseFavorite(repo *repository.FavoriteRepository) *FavoriteUseCase {
	return &FavoriteUseCase{
		repository: repo,
	}
}

func (u FavoriteUseCase) AddAndRemoveFavorite(fav *model.RequestFavorite) error {
	favorite := &model.RequestFavorite{
		Username: fav.Username,
		IdCar:    fav.IdCar,
	}

	err := u.repository.AddAndRemoveFavorite(favorite)
	return err
}

func (u FavoriteUseCase) GetFavorite(username string) (*[]model.ResponseFavorite, error) {
	result, err := u.repository.GetFavorite(username)

	return result, err
}
