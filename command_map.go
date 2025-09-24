package main

import (
	"fmt"

	"github.com/riyadvr/pokedexcli/pokeapi"
)

func commandMap(cfg *pokeapi.Config, userInput string) error {
	url := ""
	if cfg.Next != nil {
		url = *cfg.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	location, err := cfg.Client.GetLocation(url)
	if err != nil {
		return err
	}
	if location.Next != nil {
		cfg.Next = location.Next
	}
	if location.Previous != nil {
		cfg.Previous = location.Previous
	}
	for _, name := range location.Results {
		fmt.Println(name.Name)
	}
	return nil

}

func commandMapB(cfg *pokeapi.Config, userInput string) error {
	url := *cfg.Previous
	location, err := cfg.Client.GetLocation(url)
	if err != nil {
		return err
	}
	if location.Next != nil {
		cfg.Next = location.Next
	}
	if location.Previous != nil {
		cfg.Previous = location.Previous
	}
	for _, name := range location.Results {
		fmt.Println(name.Name)
	}
	return nil
}
