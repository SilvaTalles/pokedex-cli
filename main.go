package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commandRegistry = getRegistry()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		firstWord := words[0]

		command, exits := commandRegistry[firstWord]
		if !exits {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.callback(); err != nil {
			fmt.Println("Error: ", err)
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("Invalid input %s", err)
		}
	}
}
