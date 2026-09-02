package main

import (
	"time"

	"github.com/SilvaTalles/pokedex-cli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		commandList:   getRegistry(),
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
