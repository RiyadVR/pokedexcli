package main

import (
	"fmt"

	"github.com/riyadvr/pokedexcli/pokeapi"
)

func commandPokedex(cfg *pokeapi.Config, userInput string) error {
	for key := range pokemonCollection {
		fmt.Printf(" - %s\n", key)
	}
	return nil
}
