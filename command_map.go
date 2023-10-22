package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/jming514/pokedex-go/internal/pokeapi"
)

func (cfg *config) commandMap(_ ...string) error {
	queryUrl := locationUrl
	if cfg.nextLocationURL != nil {
		queryUrl = *cfg.nextLocationURL
	}
	cachedData, bool := cfg.pokeapiClient.C.Get(queryUrl)
	data := pokeapi.PokedexLocation{}

	if !bool {
		res, err := http.Get(queryUrl)
		if err != nil {
			log.Printf("error fetching location data: %v\n", err)
			return err
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Printf("error reading res.body: %v\n", err)
			return err
		}
		defer res.Body.Close()

		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Printf("error unmarshalling location data: %v\n", err)
			return err
		}

		cfg.pokeapiClient.C.Add(queryUrl, body)
	} else {
		err := json.Unmarshal(cachedData, &data)
		if err != nil {
			log.Printf("error unmarshalling location data: %v\n", err)
			return err
		}
	}

	printResults(data.Results)

	cfg.nextLocationURL = &data.Next
	cfg.prevLocationURL = &data.Previous

	// save to cache
	return nil
}
