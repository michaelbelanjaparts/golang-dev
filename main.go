package main

import (
	"golang-dev/handler"
	"golang-dev/utils"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	server := utils.Setup()
	db := server.InitDB()
	handlers := handler.SetupHandler(db)
	route := Router(handlers)

	log.Printf("[SERVER] starting in port : %v", os.Getenv("SERVER_PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), route); err != nil {
		panic(err)
	}
}

func Router(handlers handler.Handlers) http.Handler {
	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Get("/", handlers.GetAllHandlers)
		r.Get("/read", handlers.GetAllHandlers)
		r.Get("/create", handlers.InsertDataHandler)
		r.Get("/update", handlers.EditDataHandler)
		r.Get("/delete", handlers.DeleteDataHandler)
	})
	return router
}
