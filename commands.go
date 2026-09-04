package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config, arg string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, arg string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, com := range cfg.commandList {
		fmt.Printf("%v: %v\n", com.name, com.description)
	}
	return nil
}

func commandMapf(cfg *config, arg string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.prevLocationURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println("-", loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, arg string) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.prevLocationURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println("-", loc.Name)
	}
	return nil
}

func commandExplore(cfg *config, arg string) error {
	pokemonResp, err := cfg.pokeapiClient.ListPokemon(arg)
	if err != nil {
		return err
	}
	fmt.Println("==== Exploring", arg, "====")

	for _, poke := range pokemonResp.PokemonEncounters {
		fmt.Println("-", poke.Pokemon.Name)
	}
	return nil
}
