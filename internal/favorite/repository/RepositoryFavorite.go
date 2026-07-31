package repository

import (
	"database/sql"
	"rental_car/internal/favorite/model"
)

type FavoriteRepository struct {
	db *sql.DB
}

func NewRepositoryFavorite(db *sql.DB) *FavoriteRepository {
	return &FavoriteRepository{
		db: db,
	}
}

func (r FavoriteRepository) Favorite(fav *model.RequestFavorite) error {
	query := `SELECT EXISTS(
    SELECT 1 FROM user_favorite WHERE username = ? AND id_car = ?)`

	var isExist bool
	err := r.db.QueryRow(query, fav.Username, fav.IdCar).Scan(&isExist)
	if err != nil {
		return err
	}

	if isExist {
		query := `DELETE FROM user_favorite where username = ? AND id_car = ?`
		_, err := r.db.Exec(query, fav.Username, fav.IdCar)
		if err != nil {
			return err
		}
	} else {
		query := `INSERT INTO user_favorite (username, id_car) VALUES (?,?)`
		_, err := r.db.Exec(query, fav.Username, fav.IdCar)
		if err != nil {
			return err
		}
	}

	return err
}
