package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemons(url string) (Pokemon, error) {
	var pokemon Pokemon

	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &pokemon); err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, nil
	}

	c.cache.Add(url, data)

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, nil
	}

	return pokemon, nil
}
