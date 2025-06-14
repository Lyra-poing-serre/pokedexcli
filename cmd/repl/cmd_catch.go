package repl

import (
	"fmt"
	"math/rand"

	"github.com/Lyra-poing-serre/pokedexcli/internal/pokeapi"
	"github.com/Lyra-poing-serre/pokedexcli/internal/pokecache"
)

func commandCatch(c *Config, pokemonName string) error {
	if _, exist := c.Pokedex.Get(pokemonName); exist {
		return fmt.Errorf("%s already caught.", pokemonName)
	}

	pokemonResp, err := c.HttpClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonResp.Name)

	if catchLogic(c, pokemonResp) {
		fmt.Printf("%s was caught!\n", pokemonResp.Name)
		c.Pokedex.Add(newPokemon(pokemonResp))
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemonResp.Name)
	}
	return nil
}

func catchLogic(c *Config, pokemonResp pokeapi.PokemonResponse) bool {
	nbTries, exist := c.HttpClient.Cache.Get(pokemonResp.Name)
	userChance := rand.Intn(pokemonResp.BaseExperience * 2) // catch si au moins chance >= de la base exp
	triesBuffer := make([]byte, 1)
	if userChance < pokemonResp.BaseExperience {
		if !exist {
			triesBuffer[0] = byte(1) // + 10% base exp par try
			c.HttpClient.Cache.Add(pokemonResp.Name, triesBuffer)
			return false
		}
		nbTries := int(nbTries[0])
		bonus := pokemonResp.BaseExperience * ((nbTries * 10) / 100) // part de l'experience de base
		if (userChance + bonus) < pokemonResp.BaseExperience {
			triesBuffer[0] = byte(nbTries + 1)
			return false
		}
	}
	return true
}

func newPokemon(pokemon pokeapi.PokemonResponse) pokecache.Pokemon {
	pkmTypes := make([]string, len(pokemon.Types))
	for i, new := range pokemon.Types {
		pkmTypes[i] = new.Type.Name
	}

	pkmStats := make([]pokecache.StatsSummary, len(pokemon.Stats))
	for i, new := range pokemon.Stats {
		pkmStats[i] = pokecache.StatsSummary{
			Name:     new.Stat.Name,
			BaseStat: new.BaseStat,
		}
	}

	return pokecache.Pokemon{
		Id:             pokemon.ID,
		Name:           pokemon.Name,
		BaseExperience: pokemon.BaseExperience,
		Height:         pokemon.Height,
		Weight:         pokemon.Weight,
		Stats:          pkmStats,
		Types:          pkmTypes,
	}
}
