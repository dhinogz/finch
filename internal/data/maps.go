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
	Start     Place
	End       Place
	Waypoints []Place
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
			Polyline string `json:"polyline"`
		} `json:"sections"`
	} `json:"routes"`
}

func (mm *MapModel) GetDefaultRoute() (*Route, error) {

	route := &Route{
		Start: Place{
			Name: "Berlin1",
			Lat:  52.522297,
			Lng:  13.353296,
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

func getBox(avoid []Place) string {
	// &avoid[areas]=polygon:34.07508,-117.94758;33.89025,-117.73334;33.89759,-118.14015|polygon:54.073334,-118.027496;53.888504,-117.813255;53.895847,-118.22007
	// for loop for every place in avoid
	var avoidBox string

	//check if avoid has only one place
	if len(avoid) == 1 {
		box1 := Place{Lat: avoid[0].Lat + 0.00300, Lng: avoid[0].Lng - 0.00300}
		box2 := Place{Lat: avoid[0].Lat - 0.00300, Lng: avoid[0].Lng + 0.00300}
		avoidBox = fmt.Sprintf("bbox:%f,%f,%f,%f", box1.Lng, box1.Lat, box2.Lng, box2.Lat)
	} else {
		for _, place := range avoid {
			box1 := Place{Lat: place.Lat + 0.00300, Lng: place.Lng - 0.00300}
			box2 := Place{Lat: place.Lat - 0.00300, Lng: place.Lng + 0.00300}
			avoidBox = fmt.Sprintf("bbox:%f,%f,%f,%f", box1.Lng, box1.Lat, box2.Lng, box2.Lat)
			// print avoidBox
			fmt.Println(avoidBox)
		}
	}
	return avoidBox
}

func (mm *MapModel) CalcRoute() (*Route, error) {
	apiKey := "6v2fXdPA23DAqav7sAqa8JRo7xfi-KlV6hySAwOkKbM"

	// generate box to avoid

	// TODO: implement this!!!
	// avoid, err := mm.GetDangerousArea()
	// if err != nil {
	// 	return nil, err
	// }
	avoid := []Place{
		{
			Lat: 52.514706,
			Lng: 13.369239,
		},
	}

	box := getBox(avoid)

	//avoid := "[areas]=" + box

	originStr := "52.522297,13.353296"
	destinationStr := "52.508309,13.355633"
	start := Place{
		Lat: 52.522297,
		Lng: 13.353296,
	}
	destination := Place{
		Lat: 52.508309,
		Lng: 13.355633,
	}
	route := Route{
		Start: start,
		End:   destination,
	}
	avoidStr := "[areas]=" + box
	// avoidStr := fmt.Sprintf("[areas]=bbox:%s", avoid)
	apiUrl := "https://router.hereapi.com/v8/routes?" +
		"origin=" + originStr +
		"&destination=" + destinationStr +
		"&transportMode=car" +
		"&avoid" + avoidStr +
		"&return=polyline" +
		"&apikey=" + apiKey

	fmt.Println(apiUrl)
	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var routeResponse RouteResponse
	err = json.Unmarshal(responseBody, &routeResponse)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return nil, err
	}

	var polyline string
	for _, route := range routeResponse.Routes {
		for _, section := range route.Sections {
			fmt.Println("Polyline:", section.Polyline)
			polyline = section.Polyline
		}
	}

	dec, error := flexpolyline.Decode(polyline)
	if error != nil {
		fmt.Println("Error decoding polyline:", error)
		return nil, err
	}
	var waypoints []Place
	// get the lat and lng values from dec.Coordinates() and place in a matrix
	// define a counter for the for loop
	var c, skip, i int
	for c < len(dec.Coordinates()) {
		skip = (len(dec.Coordinates()) / 6) + 1
		// check if skip is divisible by 25
		if skip <= 1 {
			c += 1
		} else {
			c += skip
		}
		if c >= len(dec.Coordinates()) {
			break
		}
		i += 1
		p := Place{Lat: dec.Coordinates()[c].Lat, Lng: dec.Coordinates()[c].Lng}
		waypoints = append(waypoints, p)
	}
	fmt.Println(len(waypoints))
	route.Waypoints = waypoints

	return &route, nil
}
