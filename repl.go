package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/SilvaTalles/pokedex-cli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	commandList     map[string]cliCommand
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		firstWord := words[0]

		// getting the arg if exists
		arg := ""
		if len(words) > 1 {
			arg = words[1]
		}

		command, exists := cfg.commandList[firstWord]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.callback(cfg, arg); err != nil {
			fmt.Println("Error: ", err)
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Invalid input %s", err)
		}
	}
}

func getRegistry() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 map locations in the Pokemon world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 map locations in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays all the pokemons in the given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catches a pokemon",
			callback:    commandCatch,
		},
		"open": {
			name:        "open",
			description: "Open the Pokedex for cought pokemons",
			callback:    commandOpen,
		},
	}
}

func cleanInput(text string) []string {
	var result []string
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	i := strings.Index(text, " ")
	for i > -1 {
		result = append(result, text[:i])
		text = text[i+1:]
		i = strings.Index(text, " ")
	}
	result = append(result, text)
	return result
}
