package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jkk290/pokedexcli/internal/pokecache"
)

type Client struct {
	http  *http.Client
	cache *pokecache.Cache
}

func NewClient(interval time.Duration) *Client {
	return &Client{
		http:  &http.Client{},
		cache: pokecache.NewCache(interval),
	}
}

func (c *Client) GetLocations(url string) (Location, error) {
	if cachedData, exists := c.cache.Get(url); exists {
		locations := Location{}
		if err := json.Unmarshal(cachedData, &locations); err != nil {
			return Location{}, fmt.Errorf("error converting json from cache: %v", err)
		}
		fmt.Println("Successfully read from cache")
		return locations, nil
	}

	res, err := c.http.Get(url)
	if err != nil {
		return Location{}, fmt.Errorf("error getting locations: %v", err)
	}
	defer res.Body.Close()

	resBodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, fmt.Errorf("error converting response body to bytes: %v", err)
	}
	c.cache.Add(url, resBodyBytes)
	locations := Location{}

	if err := json.Unmarshal(resBodyBytes, &locations); err != nil {
		return Location{}, fmt.Errorf("error converting json: %v", err)
	}
	return locations, nil
}

type Location struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name *string `json:"name"`
		URL  *string `json:"url"`
	} `json:"results"`
}
