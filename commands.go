package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"slices"
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

func commandCatch(cfg *config, arg string) error {
	catchResp, err := cfg.pokeapiClient.GetPokemonInfo(arg)
	if err != nil {
		return err
	}

	chance := 100 - (catchResp.BaseExperience / 5)
	hit := rand.Intn(100)

	fmt.Printf("-> Throwing a Pokeball at %v...\n", arg)
	if hit <= chance {
		cfg.pokeapiClient.RegisterPokemon(arg, catchResp)
		fmt.Printf("**%v was caught!**\n", arg)
	} else {
		fmt.Printf("**%v escaped!**\n", arg)
	}

	return nil
}

func commandOpen(cfg *config, arg string) error {
	list := cfg.pokeapiClient.GetPokedexList()
	if len(list) == 0 {
		fmt.Println("You don't have any Pokemon yet!")
	}
	for _, poke := range list {
		fmt.Println("-", poke)
	}
	return nil
}

func commandInspect(cfg *config, arg string) error {
	infoResp, err := cfg.pokeapiClient.GetPokemonInfo(arg)
	if err != nil {
		return err
	}

	list := cfg.pokeapiClient.GetPokedexList()
	if !slices.Contains(list, arg) {
		fmt.Println("You haven't caught this Pokemon")
		return nil
	}

	fmt.Println("#Name:", infoResp.Name)
	fmt.Println("#Height:", infoResp.Height)
	fmt.Println("#Weight:", infoResp.Weight)
	fmt.Println("#Stats:")
	fmt.Println("   -hp:", infoResp.Stats[0].BaseStat)
	fmt.Println("   -attack:", infoResp.Stats[1].BaseStat)
	fmt.Println("   -defense:", infoResp.Stats[2].BaseStat)
	fmt.Println("   -special-attack:", infoResp.Stats[3].BaseStat)
	fmt.Println("   -special-defense:", infoResp.Stats[4].BaseStat)
	fmt.Println("   -speed", infoResp.Stats[5].BaseStat)
	fmt.Println("#Types:")
	for _, pokeType := range infoResp.Types {
		fmt.Println("   -", pokeType.Type.Name)
	}

	return nil
}
