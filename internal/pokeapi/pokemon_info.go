package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type PokemonInfo struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func (c *Client) GetPokemonInfo(arg string) (PokemonInfo, error) {
	url := baseURL + "/pokemon/" + arg

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfo{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	pokeInfo := PokemonInfo{}
	err = json.Unmarshal(data, &pokeInfo)
	if err != nil {
		return PokemonInfo{}, err
	}

	return pokeInfo, nil
}
