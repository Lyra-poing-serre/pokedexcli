package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (c *Client) GetPokemon(pokemonName string) (pokemonResponse PokemonResponse, err error) {
	if pokemonName == "" {
		return PokemonResponse{}, errors.New("Empty pokemonName...")
	}
	body, err := c.getRequest(fmt.Sprint(baseUrl, "/pokemon/", pokemonName))
	if err != nil {
		return PokemonResponse{}, fmt.Errorf("GET request error with url: %s got this error: %w", pokemonName, err)
	}

	if err = json.Unmarshal(body, &pokemonResponse); err != nil {
		return PokemonResponse{}, fmt.Errorf("Failed to decode json, got err: %w\n", err)
	}
	return pokemonResponse, nil
}
