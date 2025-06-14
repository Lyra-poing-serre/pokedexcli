package main

import (
	"time"

	"github.com/Lyra-poing-serre/pokedexcli/cmd/repl"
	"github.com/Lyra-poing-serre/pokedexcli/internal/pokeapi"
	"github.com/Lyra-poing-serre/pokedexcli/internal/pokecache"
)

func main() {
	repl.StartRepl(
		&repl.Config{
			HttpClient: pokeapi.NewClient(30*time.Second, 5*time.Minute),
			Pokedex:    pokecache.Pokedex{Pokedex: make(map[string]pokecache.Pokemon)},
		})
}
