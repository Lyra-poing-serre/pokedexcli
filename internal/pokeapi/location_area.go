package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (c *Client) GetLocationArea(locationAreaUrl *string) (locationAreaResponse LocationAreaResponse, err error) {
	url := baseUrl + "/location-area"
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

type AreaResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}
