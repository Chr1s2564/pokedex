package pokeapi

import (
	"pokedex/internal/pokecache"
	"time"
)

type Client struct {
	cache *pokecache.Cache
}

func NewClient(interval time.Duration) *Client {
	return &Client{
		cache: pokecache.NewCache(interval),
	}
}
