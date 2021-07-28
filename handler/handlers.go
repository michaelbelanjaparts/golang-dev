package handler

import (
	"database/sql"
	"fmt"
	"golang-dev/model"
	"golang-dev/utils"
	"net/http"
	"net/url"

	"golang-dev/usecase"
)

type handler struct {
	db *sql.DB
}

type Handlers interface {
	GetAllHandlers(w http.ResponseWriter, r *http.Request)
	InsertDataHandler(w http.ResponseWriter, r *http.Request)
	EditDataHandler(w http.ResponseWriter, r *http.Request)
	DeleteDataHandler(w http.ResponseWriter, r *http.Request)
}

func SetupHandler(db *sql.DB) Handlers {
	return &handler{db: db}
}

func (h *handler) GetAllHandlers(w http.ResponseWriter, r *http.Request) {
	res, err := usecase.GetAllArtist(h.db)
	if err != nil {
		res_err := model.SimpleResponse{Message: fmt.Sprintf("%v", err)}
		utils.Response(w, http.StatusBadRequest, res_err)
		return
	}

	utils.Response(w, http.StatusOK, res)
	return
}

func (h *handler) InsertDataHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 8*1024*1024)

	// read file data
	image, headers, err := r.FormFile("file")
	if err != nil {
		utils.Response(w, http.StatusUnprocessableEntity, err)
		return
	}
	defer image.Close()

	err = usecase.InsertArtist(h.db, image, url.QueryEscape(headers.Filename))
	if err != nil {
		utils.Response(w, http.StatusBadRequest, err)
		return
	}

	utils.Response(w, http.StatusOK, "success")
	return
}

func (h *handler) EditDataHandler(w http.ResponseWriter, r *http.Request) {
	utils.Response(w, http.StatusOK, "res")
	return
}

func (h *handler) DeleteDataHandler(w http.ResponseWriter, r *http.Request) {
	utils.Response(w, http.StatusOK, "res")
	return
}
