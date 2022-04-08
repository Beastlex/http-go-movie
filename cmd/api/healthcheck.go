package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckerHandler(w http.ResponseWriter, r *http.Request) {
	js := `{"status": "available", "env": %q, "ver": %q}`
	js = fmt.Sprintf(js, app.config.env, version)

	// Set correct content type in header
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(js))
}
