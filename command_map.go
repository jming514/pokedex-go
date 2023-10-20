package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/jming514/pokedex-go/internal/pokeapi"
)

func (cfg *config) commandMap() error {
	// check cache
	// check config if there is a Next already
	// if there is a Next, then use it
	// if not, then use the default locationUrl
	// save Next after fetching
	// save to cache

	queryUrl := locationUrl
	if cfg.nextLocationURL != nil {
		queryUrl = *cfg.nextLocationURL
	}

	res, err := http.Get(queryUrl)
	if err != nil {
		log.Printf("error fetching location data: %v", err)
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("error reading res.body: %v", err)
		return err
	}
	defer res.Body.Close()

	data := pokeapi.PokedexLocation{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("error unmarshalling location data: %v", err)
		return err
	}

	printResults(data.Results)

	cfg.nextLocationURL = &data.Next
	cfg.prevLocationURL = &data.Previous

	// save to cache
	cfg.pokeapiClient.C.Add(queryUrl, body)
	return nil
}
