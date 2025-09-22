package pokeapi

import (
	"encoding/json"
	"net/http"
)

type Config struct {
	Next     *string
	Previous *string
}

var Cfg Config

type Location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocation(url string) (Location, error) {
	res, err := http.Get(url)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	var location Location
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&location); err != nil {
		return Location{}, err
	}

	if location.Next != nil {
		Cfg.Next = location.Next
	}

	if location.Previous != nil {
		Cfg.Previous = location.Previous
	}

	return location, nil
}
