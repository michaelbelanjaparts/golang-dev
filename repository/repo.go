package repository

import (
	"database/sql"
	"golang-dev/model"
)

func FindAll(db *sql.DB) (res []model.ArtistEntity, err error) {
	rows, err := db.Query("SELECT id, name, album, image_url, release_date, price, sample_url FROM `artist`;")
	if err != nil {
		return res, err
	}
	for rows.Next() {
		data := model.ArtistEntity{}
		err = rows.Scan(&data.ID, &data.Name, &data.Album, &data.Image, &data.Date, &data.Price, &data.SampleURL)
		if err != nil {
			return res, err
		}
		res = append(res, data)
	}
	return res, err
}

func InsertOne(db *sql.DB, req model.ArtistViewModel) (err error) {
	query := "INSERT INTO `artist` (`name`, `album`, `image_url`, `release_date`, `price`, `sample_url`) VALUES($1, $2, $3, $4, $5, $6);"
	row := db.QueryRow(query, req.Name, req.Album, req.Image, req.Date, req.Price, req.SampleURL)

	if err := row.Err(); err != nil {
		return err
	}

	return err
}

func PutOne(db *sql.DB, req model.ArtistViewModel) (err error) {
	query := "UPDATE artist SET `name` = $2, `album` = $3, `image_url` = $4, `release_date` = $5, `price` = $6, `sample_url` = $7 WHERE id = $1;"
	row := db.QueryRow(query, req.ID, req.Name, req.Album, req.Image, req.Date, req.Price, req.SampleURL)

	if err := row.Err(); err != nil {
		return err
	}

	return err
}

func DeleteOne(db *sql.DB, id int64) (err error) {
	query := "DELETE FROM additional_discounts WHERE disc_key = $1;"
	row := db.QueryRow(query, id)

	if err := row.Err(); err != nil {
		return err
	}

	return err
}
