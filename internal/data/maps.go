package data

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dhinogz/finch/pkg/flexpolyline"
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
			Name: "Teconologico de Monterrey",
			Lat:  25.650711,
			Lng:  -100.289449,
		},
		End: Place{
			Name: "Casita de la abuela",
			Lat:  25.636026,
			Lng:  -100.315968,
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

	origin := "25.650711,-100.289449"
	destination := "25.636026,-100.315968"
	//avoid := "[areas]=bbox:13.37588,52.51061,13.34226,52.51892"
	avoid := "[areas]=bbox:25.640095,-100.312389,25.654603,-100.289249"
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
	dec, error := flexpolyline.Decode("BG45z9wB_-gp_FzF8GzKwMvMgPrEzFnVjcnG3I3S_YvM7Q7QvWvM7QvHnLrYnfjIrJzZriB7LrOjDrErEzF3DzFrE7GrEzF7G3IzKrO7G_J7Q7VzZjhB7QvW3mBj1BnGrJnLnQ7BrE7BnG3DzKzKvgBrT3_BnB_E7Qn4B_O_2B3DzK_E3InGrJjNrTjIzPrJvR7BjXnBjS7BrOjDnavC3SvH74BvC7LjD7L3D3IrEvH_JrJjIvMjD3IvCjN_EzevH3IjD_ET3DnBvCnGzKvRvgBnQzU7Q3S_EzFrJzKvCjDrT7VjIrJjSnVnVrYjS_TvCvCnVrd7G_JvRrY7BnGjDzF7Q3X_E7GnLzPjIjNnGzKnGrJ3IzP_JzP7BvC7GrEvHzKvH_JjDrEzK_O8V_iBoQ7agKrO4X7pBTnGUzFwC3N4D_EsEzFwHrJoGjIoG3IkI7LwHnLkIvMwH7LsiBr2BoV3hBkIvM0F3IsEvHsJrOgKrY8V7kBgFjI8GvW4D_JA_JnGnG_JrJjI3I7GjI_OvW7G_JnG3IrEnGrJrOvMvR_EjIrE7G3DvHjDvH4IvH8LnLwWvRA3DvC7G3DTvWwR7V8QvHgFrE4DzF0FvR8Q_JgKjIsJ3IoL_E8G7BwC_JgP_JgPjDsEvH0KjIkNrJsT7GgPrJoa3IwW3D0KjD4InG8QzFwMrEsJToBnGgK3DrE3NzF_Y_ErJT")
	if error != nil {
		fmt.Println("Error decoding polyline:", error)
		return
	}
	fmt.Println(dec)

	for _, route := range routeResponse.Routes {
		for _, section := range route.Sections {
			for _, span := range section.Spans {
				for _, name := range span.Names {
					fmt.Println(name.Value)
					streetNames = append(streetNames, name.Value)
				}
			}
		}
	}

}
