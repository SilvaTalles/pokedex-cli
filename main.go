package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		words := cleanInput(input)
		firstWord := words[0]
		fmt.Printf("Your command was: %v\n", firstWord)

		if err := scanner.Err(); err != nil {
			fmt.Printf("Invalid input %s", err)
		}
	}
}
