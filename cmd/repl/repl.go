package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Lyra-poing-serre/pokedexcli/internal/pokeapi"
	"github.com/Lyra-poing-serre/pokedexcli/internal/pokecache"
)

type Config struct {
	nextUrl    *string
	prevUrl    *string
	HttpClient pokeapi.Client
	Pokedex    pokecache.Pokedex
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *Config, argName string) error
}

type command interface {
	Help() string
}

func (c cliCommand) Help() string {
	return fmt.Sprintf("%s: %s", c.name, c.description)
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the name of the 20 next areas.",
			callback:    initMap(configNext),
		},
		"mapb": {
			name:        "mapb",
			description: "Display the name of the 20 previous areas.",
			callback:    initMap(configPrevious),
		},
		"explore": {
			name:        "explore",
			description: "Explore an <area> and found pokemon.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a <pokemon> !",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a <pokemon> characteristics. He must first be caught.",
			callback:    commandInspect,
		},
	}
}

func StartRepl(conf *Config) {
	commandDict := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		res := scanner.Scan()
		if !res {
			fmt.Println("No more inputs, exiting...")
			break
		}
		inputs := cleanInput(scanner.Text())
		cmd, ok := commandDict[inputs[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback(conf, strings.Join(inputs[1:], " "))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(
		strings.ToLower(text))
}
