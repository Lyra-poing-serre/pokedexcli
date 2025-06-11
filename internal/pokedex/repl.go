package pokedexcli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		res := scanner.Scan()
		if !res {
			fmt.Println("No more inputs, exiting...")
			break
		}
		input := cleanInput(scanner.Text())
		fmt.Printf("Your command was: %s\n", input[0])
	}
}

func cleanInput(text string) []string {
	return strings.Fields(
		strings.ToLower(text))
}
