package pokeapi

import (
	"github.com/jming514/pokedex-go/internal/pokecache"
	"time"
)

type Client struct {
	C cache.Cache
}

func NewClient(interval time.Duration) Client {
	return Client{
		C: cache.NewCache(interval),
	}
}
