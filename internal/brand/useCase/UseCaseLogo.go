package useCase

import (
	"rental_car/internal/brand/model"
	"rental_car/internal/brand/repository"
)

type LogoUseCase struct {
	repository *repository.LogoRepository
}

func NewUseCaseLogo(repository *repository.LogoRepository) *LogoUseCase {
	return &LogoUseCase{
		repository: repository,
	}
}

func (c LogoUseCase) UploadLogo(urlLogo string, name string) error {
	err := c.repository.UploadLogo(urlLogo, name)
	if err != nil {
		return err
	}
	return nil
}

func (c LogoUseCase) GetAllLogo() (*[]model.ResponseLogo, error) {
	result, err := c.repository.GetAllLogo()
	if err != nil {
		return nil, err
	}

	return result, nil
}
