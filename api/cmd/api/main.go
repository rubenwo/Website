package main

import (
	"log"
	"net/http"

	"github.com/rubenwo/Website/api/internal/api"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	err := api.InitAPI("./assets/projects.json")
	if err != nil {
		log.Fatal(err)
	}
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/projects", api.ProjectsEndpoint)
	router.Get("/projects/{name}", api.ProjectEndpoint)

	log.Println("api initialized")

	http.ListenAndServe(":80", router)
}
