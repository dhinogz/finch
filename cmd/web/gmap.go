package main

import "googlemaps.github.io/maps"

type GMapClient struct {
	APIKey string
}

func openGMap(cfg config) (*maps.Client, error) {
	return maps.NewClient(maps.WithAPIKey(cfg.gmaps.apiKey))
}
