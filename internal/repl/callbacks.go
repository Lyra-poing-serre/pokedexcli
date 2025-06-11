package repl

import (
	"errors"
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func initHelp(dict *map[string]cliCommand) (func() error, error) {
	if dict == nil {
		return func() error { return nil }, errors.New("Empty command dictionnary !")
	}
	return func() error {
		fmt.Println("Welcome to the Pokedex!\nUsage:")
		fmt.Print("\n") // IDE bitching println fini avec \n
		for _, cmd := range *dict {
			fmt.Println(cmd.Help())
		}
		return nil
	}, nil

}
