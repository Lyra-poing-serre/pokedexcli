package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationArea(areaUrl *string) (areaResponse LocationAreaResponse, err error) {
	url := BaseUrl + "/location-area"
	if areaUrl != nil {
		url = *areaUrl
	}

	if url == "" {
		return LocationAreaResponse{}, errors.New("You're on the first page")
	}
	body, err := c.getRequest(url)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("GET request error with url: %s got this error: %w", url, err)
	}

	if err := json.Unmarshal(body, &areaResponse); err != nil {
		return LocationAreaResponse{}, fmt.Errorf("Failed to decode json, got err: %w\n", err)
	}

	return areaResponse, nil
}
