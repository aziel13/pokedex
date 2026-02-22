package pokeapi

import (
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

// Client -
type Client struct {
	httpClient http.Client

	cache *pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCacheEntry(timeout),
	}
}
