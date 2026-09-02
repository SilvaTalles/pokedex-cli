package main

import (
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandRegistry map[string]cliCommand

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
