package model

import "database/sql"

type SimpleResponse struct {
	Message string `json:"message"`
}

type ArtistViewModel struct {
	ID        int64  `json:"id"`
	Number    int    `json:"number"`
	Name      string `json:"name"`
	Album     string `json:"album"`
	Image     string `json:"image"`
	Date      string `json:"date"`
	SampleURL string `json:"sample_url"`
	Price     int    `json:"price"`
}

type ArtistEntity struct {
	ID        sql.NullInt64   `db:"id"`
	Name      string          `db:"name"`
	Album     string          `db:"album"`
	Image     string          `db:"image"`
	Date      string          `db:"date"`
	SampleURL string          `db:"sample_url"`
	Price     sql.NullFloat64 `db:"price"`
}

type PeekImageData struct {
	Headers interface{} `json:"headers"`
	Image   interface{} `json:"image"`
}
