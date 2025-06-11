package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type command interface {
	Help() string
}

func (c cliCommand) Help() string {
	return fmt.Sprintf("%s: %s", c.name, c.description)
}

func StartRepl() {
	commandDict := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	commandHelp, err := initHelp(&commandDict)
	if err != nil {
		fmt.Printf("Error while creating help func: %v", err)
		return
	}
	commandDict["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		res := scanner.Scan()
		if !res {
			fmt.Println("No more inputs, exiting...")
			break
		}
		inputs := cleanInput(scanner.Text())
		for _, ipt := range inputs {
			cmd, ok := commandDict[ipt]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}
			err = cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(
		strings.ToLower(text))
}
