package repl

import (
	"errors"
	"fmt"
)

func commandPokedex(c *Config, _ string) error {
	if len(c.Pokedex.Pokedex) == 0 {
		return errors.New("Empty Pokedex, catch something first.")
	}
	fmt.Println("Your pokedex:")
	for name, _ := range c.Pokedex.Pokedex {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
