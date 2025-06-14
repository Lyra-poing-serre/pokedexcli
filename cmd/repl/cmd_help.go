package repl

import (
	"fmt"
)

func commandHelp(c *Config, argName string) error {
	fmt.Println("\nWelcome to the Pokedex!\nUsage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Println(cmd.Help())
	}
	fmt.Println()
	return nil
}
