package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Create new router
	rotuter := httprouter.New()

	// Set handlers fot routes
	rotuter.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckerHandler)
	rotuter.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	rotuter.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	return rotuter
}
