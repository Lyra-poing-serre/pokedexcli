package pokecache

type StatsSummary struct {
	Name     string
	BaseStat int
}

type Pokemon struct {
	Id             int
	Name           string
	BaseExperience int
	Height         int
	Weight         int
	Stats          []StatsSummary
	Types          []string
}

type Pokedex struct {
	Pokedex map[string]Pokemon
	// pokemap map[string]Map  // Pour plus tard ?
}

func (p *Pokedex) Get(pokemonName string) (pokemon Pokemon, exist bool) {
	pokemon, exist = p.Pokedex[pokemonName]
	if !exist {
		return Pokemon{}, false
	}
	return pokemon, true
}

func (p *Pokedex) Add(pokemon Pokemon) {
	p.Pokedex[pokemon.Name] = pokemon
}
