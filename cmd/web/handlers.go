package main

import (
	"fmt"
	"net/http"
)

func (app *application) handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Health check")
}
