package main

import (
	"io"
	"log"
	"net/http"
)

func commandMapb() error {
	locationUrl := "https://pokeapi.co/api/v2/location"

	res, err := http.Get(locationUrl)
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

	log.Printf("location data: %v", string(body))
	// save to cache
	return nil
}
