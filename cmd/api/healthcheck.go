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

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "Server could not process your request", http.StatusInternalServerError)
	}
}
