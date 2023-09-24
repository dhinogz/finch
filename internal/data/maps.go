package data

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

type RouteResponse struct {
	Routes []struct {
		Sections []struct {
			Spans []struct {
				Names []struct {
					Value string `json:"value"`
				} `json:"names"`
				MaxSpeed float64 `json:"maxSpeed"`
			} `json:"spans"`
		} `json:"sections"`
	} `json:"routes"`
}

func (mm *MapModel) GetDefaultRoute() (*Route, error) {

	route := &Route{
		Start: Place{
			Name: "Berlin1",
			Lat:  52.5222969,
			Lng:  13.3532959,
		},
		End: Place{
			Name: "Berlin2",
			Lat:  52.508309,
			Lng:  13.355633,
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

func (mm *MapModel) CalcRoute() {
	apiKey := "6v2fXdPA23DAqav7sAqa8JRo7xfi-KlV6hySAwOkKbM"

	// apiUrl := "https://route.ls.hereapi.com/routing/7.2/calculateroute.json" +
	// 	"?apiKey=" + apiKey +
	// 	"&waypoint0=geo!52.5,13.4" +
	// 	"&waypoint1=geo!52.5,13.45" +
	// 	"&mode=fastest;car;traffic:disabled"

	origin := "52.522297,13.353296"
	destination := "52.508309,13.355633"
	avoid := "[areas]=bbox:13.37588,52.51061,13.34226,52.51892"

	apiUrl := "https://router.hereapi.com/v8/routes?" +
		"origin=" + origin +
		"&destination=" + destination +
		"&transportMode=car" +
		"&avoid" + avoid +
		"&spans=maxSpeed,names" +
		"&return=polyline" +
		"&apikey=" + apiKey
	fmt.Println(apiUrl)

	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println(string(responseBody))

	var routeResponse RouteResponse
	err = json.Unmarshal(responseBody, &routeResponse)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return
	}

	var streetNames []string

	for _, route := range routeResponse.Routes {
		for _, section := range route.Sections {
			for _, span := range section.Spans {
				for _, name := range span.Names {
					fmt.Println("Street Name:", name.Value)
					streetNames = append(streetNames, name.Value)
				}
			}
		}
	}

}
