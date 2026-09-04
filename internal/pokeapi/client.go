package pokeapi

import (
	"net/http"
	"time"

	"github.com/SilvaTalles/pokedex-cli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	pokedex    map[string]PokemonInfo
	httpClient http.Client
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache:   pokecache.NewCache(cacheInterval),
		pokedex: NewPokedex(),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
