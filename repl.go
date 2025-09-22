package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		userInput := scanner.Text()
		cleanInput := cleanInput(userInput)
		firstWord := cleanInput[0]
		fmt.Printf("Your command was: %s\n", firstWord)
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(strings.TrimSpace(text)))
	return words
}
