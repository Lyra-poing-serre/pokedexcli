package repl

import (
	"fmt"

	"github.com/Lyra-poing-serre/pokedexcli/internal/pokeapi"
)

type configStatus int

const (
	configPrevious configStatus = iota
	configNext
)

func initMap(status configStatus) func(c *Config) error {
	return func(c *Config) error {
		var err error
		var response pokeapi.LocationAreaResponse

		switch status {
		case configPrevious:
			response, err = c.HttpClient.GetLocationArea(c.prevUrl)
		case configNext:
			response, err = c.HttpClient.GetLocationArea(c.nextUrl)
		}

		if err != nil {
			return err
		}

		for _, area := range response.Results {
			fmt.Println(area.Name)
		}
		c.nextUrl = &response.Next
		c.prevUrl = &response.Previous
		return nil
	}
}
