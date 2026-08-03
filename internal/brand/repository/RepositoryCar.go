package repository

import (
	"database/sql"
	"rental_car/internal/brand/model"
)

type LogoRepository struct {
	db *sql.DB
}

func NewRepositoryLogo(db *sql.DB) *LogoRepository {
	return &LogoRepository{
		db: db,
	}
}

func (r LogoRepository) UploadLogo(urlLogo string, name string) error {
	query := `INSERT INTO brand (logo, name) values (?,?)`
	_, err := r.db.Exec(query, urlLogo, name)
	if err != nil {
		return err
	}

	queryDetailCar := `
        UPDATE detail_car 
        SET logo = ? 
        WHERE brand = ?
    `
	_, err = r.db.Exec(queryDetailCar, urlLogo, name)
	if err != nil {
		return err
	}

	return nil
}

func (r LogoRepository) GetAllLogo() (*[]model.ResponseLogo, error) {
	query := `SELECT id, name, logo from brand`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var brans []model.ResponseLogo

	for rows.Next() {
		var result model.ResponseLogo
		rows.Scan(&result.Id, &result.Name, &result.Logo)
		brans = append(brans, result)
	}

	return &brans, err
}
