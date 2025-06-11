package main

import (
	"time"

	"github.com/Lyra-poing-serre/pokedexcli/cmd/repl"
	"github.com/Lyra-poing-serre/pokedexcli/internal/pokeapi"
)

func main() {
	repl.StartRepl(
		&repl.Config{HttpClient: pokeapi.NewClient(30 * time.Second)})
}
