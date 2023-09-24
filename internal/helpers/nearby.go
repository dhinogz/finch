package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type NearbySearchResponse struct {
	Results []struct {
		Name             string `json:"name"`
		PlaceID          string `json:"place_id"`
		FormattedAddress string `json:"vicinity"`
		Geometry         struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
	Status string `json:"status"`
}

const nearbyBaseURL = "https://maps.googleapis.com/maps/api/place/nearbysearch/json"

func NearbySearch(keyword, apiKey string, latitude float64, longitude float64, radius int) (*NearbySearchResponse, error) {
	endpoint := fmt.Sprintf("%s?keyword=%s&location=%f,%f&radius=%d&key=%s",
		nearbyBaseURL,
		url.QueryEscape(keyword),
		latitude,
		longitude,
		radius,
		apiKey)

	resp, err := http.Get(endpoint)
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

	var response NearbySearchResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	if response.Status != "OK" {
		return nil, fmt.Errorf("API response status: %s", response.Status)
	}

	return &response, nil
}
