package main

import (
	"context"
	"net/http"
)

func (app *application) handleMap(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	route, err := app.models.Map.CalcRoute()
	if err != nil {
		app.serverError(w, err)
		return
	}

	heatmapPoints, err := app.models.Map.GetDangerousArea()
	if err != nil {
		app.serverError(w, err)
	}

	data.Route = route
	data.MapsAPI = app.config.gmaps.apiKey
	data.HeatmapPoints = heatmapPoints
	app.renderPage(w, http.StatusOK, "home.html", data)
}

type searchRouteForm struct {
	Query string `form:"query"`
}

func (app *application) handleNewMap(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	var form searchRouteForm

	err = app.formDecoder.Decode(&form, r.PostForm)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	app.infoLog.Println(form.Query)

	ctx := context.Background()
	route, err := app.models.Map.GetRoute(ctx, "Monterrey", form.Query, app.config.gmaps.apiKey, 0.0, 0.0)
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := app.newTemplateData(r)
	data.Route = route
	data.MapsAPI = app.config.gmaps.apiKey
	app.renderPage(w, http.StatusOK, "new-map.html", data)
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
