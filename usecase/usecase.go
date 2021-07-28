package usecase

import (
	"database/sql"
	"fmt"
	"golang-dev/model"
	"golang-dev/repository"
	"io"
	"mime/multipart"
	"os"
	"strconv"
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

func GetOneArtist(db *sql.DB, id int) (res model.ArtistViewModel, err error) {
	data, err := repository.FindOne(db, id)
	if err != nil {
		return res, err
	}

	res = model.ArtistViewModel{
		ID:        data.ID.Int64,
		Name:      data.Name,
		Album:     data.Album,
		Image:     data.Image,
		Date:      data.Date,
		SampleURL: data.SampleURL,
		Price:     int(data.Price.Float64),
	}

	return res, err
}

func InsertArtist(db *sql.DB, req model.ArtistViewModel, img multipart.File, imagename string, sample multipart.File, samplename string) (err error) {
	img_url := "static/images/" + imagename
	image, err := os.OpenFile(img_url, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Error OpenFile %v", err)
		return err
	}

	defer image.Close()
	io.Copy(image, img)

	sample_url := "static/samples/" + samplename
	smp, err := os.OpenFile(sample_url, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Error OpenFile %v", err)
		return err
	}

	defer smp.Close()
	io.Copy(smp, sample)

	req.Image = img_url
	req.SampleURL = sample_url

	err = repository.InsertOne(db, req)
	if err != nil {
		fmt.Printf("Error OpenFile %v", err)
		return err
	}

	return err
}

func UpdateOne(db *sql.DB, req model.ArtistViewModel) (err error) {
	err = repository.PutOne(db, req)
	if err != nil {
		return err
	}
	return err
}

func Delete(db *sql.DB, id_string string) (err error) {
	id, err := strconv.Atoi(id_string)
	if err != nil {
		return err
	}

	err = repository.DeleteOne(db, int64(id))
	if err != nil {
		return err
	}

	return err
}
