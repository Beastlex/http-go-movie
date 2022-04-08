package main

import (
	"net/http"
)

func (app *application) healthcheckerHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "available",
		"env":    app.config.env,
		"ver":    version,
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"data": data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
