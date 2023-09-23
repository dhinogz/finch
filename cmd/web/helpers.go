package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// This helper function renders HTML template with data. If you want to render a component, write it's
// name in the block parameter
func (app *application) renderPage(w http.ResponseWriter, status int, page string, data *templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)

	err := ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) renderHTMX(w http.ResponseWriter, status int, page, component string, data interface{}) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)

	err := ts.ExecuteTemplate(w, component, data)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		CurrentDate: time.Now(),
	}
}

func (app *application) readString(qs url.Values, key, defaultValue string) string {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	return s
}
