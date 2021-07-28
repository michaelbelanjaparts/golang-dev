package usecase

import (
	"database/sql"
	"fmt"
	"golang-dev/model"
	"golang-dev/repository"
	"io"
	"mime/multipart"
	"os"
)

func GetAllArtist(db *sql.DB) (res []model.ArtistViewModel, err error) {
	data, err := repository.FindAll(db)
	if err != nil {
		return res, err
	}

	for index, artist := range data {
		single := model.ArtistViewModel{
			ID:        artist.ID.Int64,
			Number:    index,
			Name:      artist.Name,
			Album:     artist.Album,
			Image:     artist.Image,
			Date:      artist.Date,
			SampleURL: artist.SampleURL,
			Price:     int(artist.Price.Float64),
		}
		res = append(res, single)
	}

	return res, err
}

func InsertArtist(db *sql.DB, img multipart.File, filename string) (err error) {
	image, err := os.OpenFile("static/images/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Error OpenFile %v", err)
		return err
	}

	defer image.Close()
	io.Copy(image, img)

	return err
}
