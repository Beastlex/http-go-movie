package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "env: %s\n", app.config.env)
	fmt.Fprintf(w, "ver: %s\n", version)
}
