package main

import (
	"fmt"
	"strings"

	"github.com/riyadvr/pokedexcli/pokeapi"
)

func commandExplore(cfg *pokeapi.Config, userInput string) error {
	//fmt.Println("explore command executing")
	exploreArea := strings.Fields(userInput)[1]
	url := "https://pokeapi.co/api/v2/location-area/" + exploreArea

	//fmt.Println(url)
	fmt.Printf("Exploring %s...\n", exploreArea)
	pokemon, err := cfg.Client.GetPokemons(url)

	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, name := range pokemon.PokemonEncounters {
		fmt.Printf("- %s\n", name.Pokemon.Name)
	}
	return nil
}
