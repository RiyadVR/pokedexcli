package main

import (
	"fmt"
	"strings"

	"github.com/riyadvr/pokedexcli/pokeapi"
)

func commandInspect(cfg *pokeapi.Config, userInput string) error {

	pokemonName := strings.Fields(userInput)[1]

	pokemonInfo, ok := pokemonCollection[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemonInfo.Name)
	fmt.Printf("Height: %d\n", pokemonInfo.Height)
	fmt.Printf("Weight: %d\n", pokemonInfo.Weight)
	fmt.Println("Stats:")
	for _, value := range pokemonInfo.Stats {
		fmt.Printf("  -%v: %v\n", value.Stat.Name, value.BaseStat)
	}
	fmt.Println("Types:")
	for _, value := range pokemonInfo.Types {
		fmt.Printf("  - %v\n", value.Type.Name)
	}

	return nil
}
