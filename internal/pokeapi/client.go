package pokeapi

import (
	"net/http"
	"time"

	"github.com/nsp5488/pokedexcli/internal/pokecache"
)

type Client struct {
	client http.Client
	cache  *pokecache.Cache
}

func NewClient(cacheInterval time.Duration) *Client {
	return &Client{
		client: http.Client{},
		cache:  pokecache.NewCache(cacheInterval),
	}
}
