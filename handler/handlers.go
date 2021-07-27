package handler

import (
	"database/sql"
	"fmt"
	"golang-dev/model"
	"golang-dev/utils"
	"net/http"

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
		utils.Response(w, http.StatusInternalServerError, res_err)
		return
	}

	utils.Response(w, http.StatusOK, res)
	return
}

func (h *handler) InsertDataHandler(w http.ResponseWriter, r *http.Request) {
	utils.Response(w, http.StatusOK, "res")
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
