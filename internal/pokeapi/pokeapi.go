package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	cache "github.com/jming514/pokedex-go/internal/pokecache"
)

type pokedexLocation struct {
	Next     string
	Previous string
	Results  []struct {
		Name string
		URL  string
	}
	Count int
}

type config struct {
	Next     *string
	Previous *string
}

var Location = &config{
	Next:     nil,
	Previous: nil,
}

func random() {
	fmt.Println("This is something random")
}

var locationUrl = "https://pokeapi.co/api/v2/location/"

func (cfg *config) GetMap() {
	apiUrl := locationUrl
	if cfg.Next != nil {
		apiUrl = *cfg.Next
	}

	val, ok := cache.MyCache.Get(apiUrl)
	if ok {
		cfg.saveLocationToCache(val, apiUrl)
	} else {
		res, err := http.Get(apiUrl)
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		cfg.saveLocationToCache(body, apiUrl)

		err = res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (cfg *config) GetMapb() {
	apiUrl := locationUrl
	if cfg.Previous == nil {
		fmt.Println("There is no previous")
		return
	}

	apiUrl = *cfg.Previous

	val, ok := cache.MyCache.Get(apiUrl)
	if ok {
		cfg.saveLocationToCache(val, apiUrl)
	} else {
		res, err := http.Get(apiUrl)
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		cfg.saveLocationToCache(body, apiUrl)

		err = res.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (cfg *config) saveLocationToCache(val []byte, key string) {
	pl := pokedexLocation{}
	err := json.Unmarshal(val, &pl)
	if err != nil {
		log.Fatal(err)
	}

	cfg.Next = &pl.Next
	cfg.Previous = &pl.Previous
	cache.MyCache.Add(key, val)

	for _, v := range pl.Results {
		fmt.Printf("%v\n", v.Name)
	}
}
