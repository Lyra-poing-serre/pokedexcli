package repl

import (
	"errors"
	"fmt"
)

func commandInspect(c *Config, pokemonName string) error {
	pokemon, exist := c.Pokedex[pokemonName]
	if !exist {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, pkm := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", pkm.Stat.Name, pkm.BaseStat)
	}
	fmt.Println("Types:")
	for _, pkm := range pokemon.Types {
		fmt.Printf("  -%s\n", pkm.Type.Name)
	}
	return nil
}
