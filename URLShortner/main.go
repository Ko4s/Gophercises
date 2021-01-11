package main

import (
	urlshortener "github/Ko4s/urlShortener/urlShortener"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	yamlPath := "urlShortener/urls.yml"

	repo, _ := urlshortener.NewInMemoryRepo(yamlPath)

	service := urlshortener.NewService(repo)
	handler := urlshortener.NewHanlder(service)
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/{[a-z]}", handler.GetPath)

	http.ListenAndServe(":5050", r)
}
