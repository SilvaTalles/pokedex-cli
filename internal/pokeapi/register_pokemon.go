package pokeapi

func NewPokedex() map[string]PokemonInfo {
	return make(map[string]PokemonInfo)
}

func (c *Client) RegisterPokemon(pokemon string, info PokemonInfo) {
	c.pokedex[pokemon] = info
}

func (c *Client) GetPokedexList() []string {
	var list []string
	for pokemon := range c.pokedex {
		list = append(list, pokemon)
	}
	return list
}
