package pokeapi

import (
	"time"

	cache "github.com/jming514/pokedex-go/internal/pokecache"
)

type Client struct {
	C cache.Cache
}

func NewClient(interval time.Duration) Client {
	return Client{
		C: cache.NewCache(interval),
	}
}
