package main

import (
	"fmt"
	"net/http"
)

func (app *application) handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Health check")
}

type reportCreateForm struct {
	Type        string `form:"report-type"`
	Description string `form:"report-description"`
	Location    string `form:"report-location"`
}

func (app *application) handleReport(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var form reportCreateForm

	err = app.formDecoder.Decode(&form, r.PostForm)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.models.Report.Insert(form.Location, 1, form.Type, form.Description)
	fmt.Println(err)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
