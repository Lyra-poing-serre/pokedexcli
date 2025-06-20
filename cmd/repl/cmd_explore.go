package repl

import (
	"fmt"
)

func commandExplore(c *Config, areaName string) error {
	response, err := c.HttpClient.GetArea(areaName)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Exploring %s...", response.Name))
	fmt.Println("Found Pokemon:")
	for _, pokemon := range response.PokemonEncounters {
		fmt.Println(fmt.Sprintf(" - %s", pokemon.Pokemon.Name))
	}
	return nil
}
