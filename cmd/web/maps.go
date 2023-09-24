package main

import (
	"context"
	"fmt"
	"net/http"
)

func (app *application) handleMap(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	route, err := app.models.Map.GetDefaultRoute()
	if err != nil {
		app.serverError(w, err)
	}
	data.Route = route
	data.MapsAPI = app.config.gmaps.apiKey
	app.renderPage(w, http.StatusOK, "home.html", data)
}

func (app *application) handleAutocomplete(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Place string
	}

	qs := r.URL.Query()
	input.Place = app.readString(qs, "place", "")

	ctx := context.Background()

	autocomplete, err := app.models.Map.GetAutocomplete(ctx, input.Place)
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := app.newTemplateData(r)
	data.Autocomplete = autocomplete
	app.renderPage(w, http.StatusOK, "places-autocomplete.html", data)
}

func (app *application) handleNewMap(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	route, err := app.models.Map.GetRoute(ctx, "Monterrey", "Saltillo")
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := app.newTemplateData(r)
	data.Route = route
	data.MapsAPI = app.config.gmaps.apiKey
	app.renderPage(w, http.StatusOK, "new-map.html", data)
}

func (app *application) handleRoute(w http.ResponseWriter, r *http.Request) {
	places, err := app.models.Map.CalcRoute()
	if err != nil {
		app.serverError(w, err)
		return
	}
	fmt.Fprint(w, places)

}
