package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(arg string) (RespShallowPokemon, error) {
	url := baseURL + "/location-area/" + arg

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := RespShallowPokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespShallowPokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	pokemonResp := RespShallowPokemon{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return RespShallowPokemon{}, err
	}

	c.cache.Add(url, data)
	return pokemonResp, nil
}
