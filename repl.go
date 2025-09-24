package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/riyadvr/pokedexcli/pokeapi"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := pokeapi.Config{
		Client: pokeapi.NewClient(10*time.Second, 5*time.Second),
	}
	pokeapi.Cfg = cfg

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		userCommand := strings.Fields(userInput)[0]
		//provideArea := strings.Fields(userInput)[1]
		command, ok := getCommands()[userCommand]
		if !ok {
			fmt.Println("Unknown Command")

		} else {
			if err := command.callback(&pokeapi.Cfg, userInput); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(strings.TrimSpace(text)))
	return words
}
