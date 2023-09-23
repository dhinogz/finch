package data

import (
	"context"
	"database/sql"

	"googlemaps.github.io/maps"
)

type Route struct {
	Start Place
	End   Place
}

type Place struct {
	Name string
	Lat  float64
	Lng  float64
}

type MapModel struct {
	DB    *sql.DB
	GMaps *maps.Client
}

func (mm *MapModel) GetDefaultRoute() (*Route, error) {

	route := &Route{
		Start: Place{
			Name: "San Francisco",
			Lat:  37.7749,
			Lng:  -122.4194,
		},
		End: Place{
			Name: "Los Angeles",
			Lat:  34.0522,
			Lng:  -118.2437,
		},
	}

	return route, nil
}

func (mm *MapModel) GetRoute(ctx context.Context, start, destination string) (*Route, error) {

	route := &Route{
		Start: Place{
			Name: "Monterrey, Nuevo Leon",
			Lat:  25.6866,
			Lng:  -100.3161,
		},
		End: Place{
			Name: "Saltillo, Coahuila",
			Lat:  25.4383,
			Lng:  -100.9737,
		},
	}

	return route, nil
}

func (mm *MapModel) GetAutocomplete(ctx context.Context, input string) ([]string, error) {
	if input == "" {
		return make([]string, 0), nil
	}

	query := maps.QueryAutocompleteRequest{
		Input: input,
	}

	ac, err := mm.GMaps.QueryAutocomplete(ctx, &query)
	if err != nil {
		return nil, err
	}

	maxResults := 5
	numResults := min(maxResults, len(ac.Predictions))

	res := make([]string, numResults)
	for i := 0; i < numResults; i++ {
		prediction := ac.Predictions[i].Terms[0].Value
		res[i] = prediction
	}

	return res, nil

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
