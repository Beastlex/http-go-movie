package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Create new router
	rotuter := httprouter.New()

	// add our hanlers
	rotuter.NotFound = http.HandlerFunc(app.notFoundResponse)
	rotuter.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponce)

	// Set handlers fot routes
	rotuter.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckerHandler)
	rotuter.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	rotuter.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)
	rotuter.HandlerFunc(http.MethodPut, "/v1/movies/:id", app.updateMovieHandler)
	rotuter.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteMovieHandler)

	return rotuter
}
