package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/riyadvr/pokedexcli/pokeapi"
)

var PokemonCollection = make(map[string]pokeapi.PokemonInfo)

// func AppendPokemonCollection(pokemonName string) {
// 	pokemonCollection[pokemonName] = pokeapi.PokemonInfo{Name: pokemonName}
// }

func commandCatch(cfg *pokeapi.Config, userInput string) error {

	pokemonName := strings.Fields(userInput)[1]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName

	pokemonInfo, err := cfg.Client.GetPokemonInfo(url)
	if err != nil {
		return err
	}

	randomNumber := rand.Intn(500)
	// fmt.Printf("random number: %d\n", randomNumber)
	// fmt.Printf("base experience: %d\n", pokemonInfo.BaseExperience)

	if randomNumber < pokemonInfo.BaseExperience {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemonName)
	PokemonCollection[pokemonName] = pokemonInfo
	return nil
}
