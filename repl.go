package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/riyadvr/pokedexcli/pokeapi"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		command, ok := getCommands()[userInput]
		if !ok {
			fmt.Println("Unknown Command")

		} else {
			if err := command.callback(&pokeapi.Cfg); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(strings.TrimSpace(text)))
	return words
}
