package pokeapi

import (
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
