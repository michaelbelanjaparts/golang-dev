package handler

import (
	"database/sql"
	"fmt"
	"golang-dev/model"
	"golang-dev/utils"
	"net/http"
	"net/url"
	"strconv"

	"golang-dev/usecase"

	"github.com/go-chi/chi"
)

type handler struct {
	db *sql.DB
}

type Handlers interface {
	GetAllHandlers(w http.ResponseWriter, r *http.Request)
	InsertDataHandler(w http.ResponseWriter, r *http.Request)
	EditDataHandler(w http.ResponseWriter, r *http.Request)
	DeleteDataHandler(w http.ResponseWriter, r *http.Request)
	GetOneHandlers(w http.ResponseWriter, r *http.Request)
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
	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		utils.Response(w, http.StatusBadRequest, model.SimpleResponse{Message: fmt.Sprintf("%v", err)})
		return
	}

	req := model.ArtistViewModel{
		Name:  r.FormValue("name"),
		Album: r.FormValue("album"),
		Date:  r.FormValue("date"),
		Price: price,
	}

	r.Body = http.MaxBytesReader(w, r.Body, 8*1024*1024)

	image, image_header, err := r.FormFile("image")
	if err != nil {
		utils.Response(w, http.StatusUnprocessableEntity, model.SimpleResponse{Message: fmt.Sprintf("%v", err)})
		return
	}
	defer image.Close()

	sample, sample_headers, err := r.FormFile("sample")
	if err != nil {
		utils.Response(w, http.StatusUnprocessableEntity, model.SimpleResponse{Message: fmt.Sprintf("%v", err)})
		return
	}
	defer image.Close()

	err = usecase.InsertArtist(
		h.db,
		req,
		image,
		url.QueryEscape(image_header.Filename),
		sample,
		url.QueryEscape(sample_headers.Filename),
	)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, model.SimpleResponse{Message: fmt.Sprintf("%v", err)})
		return
	}

	utils.Response(w, http.StatusOK, model.SimpleResponse{Message: "success"})
	return
}

func (h *handler) EditDataHandler(w http.ResponseWriter, r *http.Request) {
	var (
		price int
		err   error
	)
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if r.FormValue("price") != "" {
		price, err = strconv.Atoi(r.FormValue("price"))
		if err != nil {
			utils.Response(w, http.StatusBadRequest, model.SimpleResponse{Message: fmt.Sprintf("%v", err)})
			return
		}
	}

	// get req data
	req := model.ArtistViewModel{
		ID:    int64(id),
		Name:  r.FormValue("name"),
		Album: r.FormValue("album"),
		Date:  r.FormValue("date"),
		Price: price,
	}

	// get db data
	res, err := usecase.GetOneArtist(h.db, id)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, model.SimpleResponse{Message: fmt.Sprintf("%v", err)})
		return
	}

	if !utils.CompareString(req.Name, res.Name) {
		res.Name = req.Name
	}

	if !utils.CompareString(req.Album, res.Album) {
		res.Album = req.Album
	}

	if !utils.CompareString(req.Date, res.Date) {
		res.Date = req.Date
	}

	if !utils.CompareInt(req.Price, res.Price) {
		res.Price = req.Price
	}

	err = usecase.UpdateOne(h.db, res)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, model.SimpleResponse{Message: fmt.Sprintf("%v", err)})
		return
	}

	// if req data != current data, update changed data
	utils.Response(w, http.StatusOK, model.SimpleResponse{Message: "success"})
	return
}

func (h *handler) DeleteDataHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	err := usecase.Delete(h.db, id)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, model.SimpleResponse{Message: fmt.Sprintf("%v", err)})
		return
	}

	utils.Response(w, http.StatusOK, model.SimpleResponse{Message: "success"})
	return
}

func (h *handler) GetOneHandlers(w http.ResponseWriter, r *http.Request) {
	string_id := chi.URLParam(r, "id")
	if string_id == "" {
		utils.Response(w, http.StatusBadRequest, model.SimpleResponse{Message: "params is required"})
		return
	}

	id, err := strconv.Atoi(string_id)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, model.SimpleResponse{Message: "params is required"})
		return
	}

	res, err := usecase.GetOneArtist(h.db, id)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, model.SimpleResponse{Message: "params is required"})
		return
	}

	utils.Response(w, http.StatusOK, res)
	return
}
