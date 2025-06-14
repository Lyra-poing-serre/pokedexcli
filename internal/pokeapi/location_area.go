package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (c *Client) GetLocationArea(locationAreaUrl *string) (locationAreaResponse LocationAreaResponse, err error) {
	url := baseUrl + locationAreaEndpoint
	if locationAreaUrl != nil {
		url = *locationAreaUrl
	}

	if url == "" {
		return LocationAreaResponse{}, errors.New("You're on the first page")
	}
	body, err := c.getRequest(url)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("GET request error with url: %s got this error: %w", url, err)
	}

	if err := json.Unmarshal(body, &locationAreaResponse); err != nil {
		return LocationAreaResponse{}, fmt.Errorf("Failed to decode json, got err: %w\n", err)
	}

	return locationAreaResponse, nil
}

type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
