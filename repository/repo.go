package repository

import (
	"database/sql"
	"golang-dev/model"
	"os"
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

func FindOne(db *sql.DB, id int) (res model.ArtistEntity, err error) {
	rows, err := db.Query("SELECT id, name, album, image_url, release_date, price, sample_url FROM `artist` WHERE id = ?;", id)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&res.ID, &res.Name, &res.Album, &res.Image, &res.Date, &res.Price, &res.SampleURL)
		if err != nil {
			return res, err
		}
	}
	return res, err
}

func InsertOne(db *sql.DB, req model.ArtistViewModel) (err error) {
	query := "INSERT INTO `artist` (`name`, `album`, `image_url`, `release_date`, `price`, `sample_url`) VALUES(?, ?, ?, ?, ?, ?);"
	row := db.QueryRow(query, req.Name, req.Album, os.Getenv("SERVER_HOST")+req.Image, req.Date, req.Price, os.Getenv("SERVER_HOST")+req.SampleURL)

	if err := row.Err(); err != nil {
		return err
	}

	return err
}

func PutOne(db *sql.DB, req model.ArtistViewModel) (err error) {
	query := "UPDATE artist SET `name` = ?, `album` = ?, `image_url` = ?, `release_date` = ?, `price` = ?, `sample_url` = ? WHERE id = ?;"
	row := db.QueryRow(query, req.Name, req.Album, req.Image, req.Date, req.Price, req.SampleURL, req.ID)

	if err := row.Err(); err != nil {
		return err
	}

	return err
}

func DeleteOne(db *sql.DB, id int64) (err error) {
	query := "DELETE FROM artist WHERE id = ?;"
	row := db.QueryRow(query, id)

	if err := row.Err(); err != nil {
		return err
	}

	return err
}
