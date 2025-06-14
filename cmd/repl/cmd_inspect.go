package repl

import (
	"errors"
	"fmt"
)

func commandInspect(c *Config, pokemonName string) error {
	pokemon, exist := c.Pokedex.Get(pokemonName)
	if !exist {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, name := range pokemon.Types {
		fmt.Printf("  -%s\n", name)
	}
	return nil
}
