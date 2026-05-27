package pokeapi

import (
	"net/http"
	"time"

	pokecache "github.com/piitah/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	HTTPClient *http.Client
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		cache: pokecache.NewCache(timeout),
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
	}
}
