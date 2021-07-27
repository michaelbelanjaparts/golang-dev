package usecase

import (
	"database/sql"
	"golang-dev/model"
	"golang-dev/repository"
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
