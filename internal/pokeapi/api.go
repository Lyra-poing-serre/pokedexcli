package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getRequest(url string) (body []byte, err error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("Error while requesting PokeAPI : %w", err)
	}
	body, err = io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("Request failed with status code: %d and body:\n %s\n", res.StatusCode, body)
	}
	if err != nil {
		return []byte{}, fmt.Errorf("Failed to read data: %w and body:\n %s", err, body)
	}
	return body, nil
}

func GetLocationArea(url string) (nextUrl, previousUrl string, areaList []string, err error) {
	body, err := getRequest(url)
	if err != nil {
		return "", "", []string{}, fmt.Errorf("GET request url: %s\ngot an error: %w", url, err)
	}

	var response LocationArea
	if err := json.Unmarshal(body, &response); err != nil {
		return "", "", []string{}, fmt.Errorf("Failed to decode json, got err: %w\n", err)
	}
	var output []string
	for _, item := range response.Results {
		output = append(output, item.Name)
	}
	return response.Next, response.Previous, output, nil
}
