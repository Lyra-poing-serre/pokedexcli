package repl

import (
	"errors"
	"fmt"
	"os"

	"github.com/Lyra-poing-serre/pokedexcli/internal/pokeapi"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func initHelpCommand(dict *map[string]cliCommand) (commandHelp func() error, err error) {
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

func initMap(config configStatus, mapConfig *config) (command func() error, err error) {
	if config > 2 {
		return func() error { return nil }, errors.New("No configStatus selected")
	} else if mapConfig.previous == "" || mapConfig.next == "" {
		return func() error { return nil }, errors.New("No URL config selected")
	}

	command = func() error {
		var targetUrl string
		switch config {
		case configPrevious:
			targetUrl = mapConfig.previous
		case configNext:
			targetUrl = mapConfig.next
		}
		next, previous, areaList, err := pokeapi.GetLocationArea(targetUrl)
		if err != nil {
			return err
		} else if len(areaList) == 0 {
			return fmt.Errorf("Empty list returned from %s", targetUrl)
		}

		for _, area := range areaList {
			fmt.Println(area)
		}
		mapConfig.next = next
		mapConfig.previous = previous
		return nil
	}
	return command, nil
}
