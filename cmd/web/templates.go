package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/dhinogz/finch/internal/data"
)

type templateData struct {
	CurrentDate  time.Time
	Form         any
	Route        *data.Route
	MapsAPI      string
	Autocomplete []string
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// Cache individual dynamic components for direct rendering using HTMX
	partials, err := filepath.Glob("./ui/html/partials/*.html")
	if err != nil {
		return nil, err
	}

	for _, d := range partials {
		name := filepath.Base(d)

		ts, err := template.New(name).Funcs(functions).ParseFiles(d)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	// Cache full pages with their components
	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, p := range pages {
		name := filepath.Base(p)

		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/components/*.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(p)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
