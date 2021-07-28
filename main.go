package main

import (
	"golang-dev/handler"
	"golang-dev/utils"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

func main() {
	server := utils.Setup()
	db := server.InitDB()
	handlers := handler.SetupHandler(db)
	route := Router(handlers)

	log.Printf("[SERVER] starting in port : %v", os.Getenv("SERVER_PORT"))
	http.Handle("/", cors.AllowAll().Handler(route))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), nil); err != nil {
		panic(err)
	}
}

func Router(handlers handler.Handlers) http.Handler {
	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Get("/", handlers.GetAllHandlers)
		r.Get("/read", handlers.GetAllHandlers)
		r.Get("/read/{id}", handlers.GetOneHandlers)
		r.Post("/create", handlers.InsertDataHandler)
		r.Put("/update", handlers.EditDataHandler)
		r.Delete("/delete", handlers.DeleteDataHandler)
	})
	return router
}
