package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(url string) (PokemonInfo, error) {
	var pokemonInfo PokemonInfo

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	if err := json.Unmarshal(data, &pokemonInfo); err != nil {
		return PokemonInfo{}, err
	}

	return pokemonInfo, nil
}
