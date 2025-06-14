package repl

import (
	"fmt"
	"math/rand"

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
	pokemon := pokecache.Pokemon{
		Id:             pokemonResp.ID,
		Name:           pokemonResp.Name,
		BaseExperience: pokemonResp.BaseExperience,
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	nbTries, exist := c.HttpClient.Cache.Get(pokemon.Name)
	userChance := rand.Intn(pokemon.BaseExperience * 2) // catch si au moins chance >= de la base exp
	triesBuffer := make([]byte, 1)

	if userChance < pokemon.BaseExperience {
		err = fmt.Errorf("%s escaped!", pokemon.Name)
		if !exist {
			triesBuffer[0] = byte(1) // + 10% base exp par try
			c.HttpClient.Cache.Add(pokemon.Name, triesBuffer)
			return err
		}

		nbTries := int(nbTries[0])
		bonus := pokemon.BaseExperience * ((nbTries * 10) / 100) // part de l'experience de base
		if (userChance + bonus) < pokemon.BaseExperience {
			triesBuffer[0] = byte(nbTries + 1)
			return err
		}
	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	c.Pokedex.Add(pokemon)
	return nil
}
