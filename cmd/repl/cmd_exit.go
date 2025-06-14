package repl

import (
	"fmt"
	"os"
)

func commandExit(c *Config, argName string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
