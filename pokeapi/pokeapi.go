package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/riyadvr/pokedexcli/pokecache"
)

type Config struct {
	Next     *string
	Previous *string
	Client   Client
}

var Cfg Config

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

type Location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache:      pokecache.NewCache(cacheInterval),
		httpClient: http.Client{Timeout: timeout},
	}
}

func (c *Client) GetLocation(url string) (Location, error) {
	var location Location

	if val, ok := c.cache.Get(url); ok {
		if err := json.Unmarshal(val, &location); err != nil {
			return Location{}, err
		}
		return location, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)

	if err := json.Unmarshal(dat, &location); err != nil {
		return Location{}, err
	}
	return location, nil
}
