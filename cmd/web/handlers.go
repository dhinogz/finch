package main

import (
	"fmt"
	"net/http"
)

func (app *application) handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Health check")
}

func (app *application) handleReport(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Form.Get("report-type"))
	fmt.Println(r.Form.Get("report-description"))
}
