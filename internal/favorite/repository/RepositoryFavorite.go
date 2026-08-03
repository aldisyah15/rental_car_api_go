package repository

import (
	"database/sql"
	"encoding/json"
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

func (r FavoriteRepository) AddAndRemoveFavorite(fav *model.RequestFavorite) error {
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

func (r FavoriteRepository) GetFavorite(username string) (*[]model.ResponseFavorite, error) {
	query := `SELECT
    username, 
    id_car,
    name,
    brand,
    rental_price,
    image,
    horse_power,
    gear,
    description,
    seat,
   IFNULL(logo, ''),
    stock,
    DATE_FORMAT(create_at, '%Y-%m-%d %H:%i:%s')
	FROM user_favorite
	INNER JOIN detail_car ON id_car = id
	where username = ?`

	rows, err := r.db.Query(query, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var AllFavorite []model.ResponseFavorite
	for rows.Next() {
		var images []byte
		var fav model.ResponseFavorite
		err := rows.Scan(&fav.Username, &fav.IdCar, &fav.Name, &fav.Brand, &fav.RentalPrice, &images, &fav.Horsepower, &fav.Gear, &fav.Description, &fav.Seat, &fav.Logo, &fav.Stock, &fav.CreateAt)
		if err != nil {
			return nil, err
		}

		if len(images) > 0 {
			_ = json.Unmarshal(images, &fav.Images)
		}

		AllFavorite = append(AllFavorite, fav)

	}

	return &AllFavorite, nil
}
