package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PlaceAPIResponse struct {
	Candidates []struct {
		FormattedAddress string  `json:"formatted_address"`
		Name             string  `json:"name"`
		Rating           float32 `json:"rating"`
		OpeningHours     struct {
			OpenNow bool `json:"open_now"`
		} `json:"opening_hours"`
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"candidates"`
	Status string `json:"status"`
}

const baseURL = "https://maps.googleapis.com/maps/api/place/findplacefromtext/json?fields=formatted_address,name,rating,opening_hours,geometry"

func FindPlaceFromText(destination, apiKey string) (*PlaceAPIResponse, error) {
	url := fmt.Sprintf("%s&input=%s&inputtype=textquery&key=%s", baseURL, destination, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response PlaceAPIResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.Status != "OK" {
		return nil, fmt.Errorf("API response status: %s", response.Status)
	}

	return &response, nil
}
