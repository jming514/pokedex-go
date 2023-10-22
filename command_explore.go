package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type exploreResponse struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
	Id int `json:"id"`
}

func (cfg *config) commandExplore(args ...string) error {
	area := args[0]

	if area == "" {
		return errors.New("no area specified")
	}

	exploreUrl := locationUrl + "-area/" + area + "-area"

	cachedData, ok := cfg.pokeapiClient.C.Get(exploreUrl)
	fmt.Println(ok)
	data := exploreResponse{}

	if !ok {
		res, err := http.Get(exploreUrl)
		if err != nil {
			log.Println("error fetching location data:", err)
			return err
		}

		if strings.Contains(res.Status, "404") {
			log.Println("error location not found:", err)
			return errors.New(res.Status)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Println("error reading res.body:", err)
			return err
		}
		defer res.Body.Close()

		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Println("error unmarshalling location data:", err)
			return err
		}
		fmt.Println("inside http section")

		cfg.pokeapiClient.C.Add(exploreUrl, body)
	} else {
		err := json.Unmarshal(cachedData, &data)
		if err != nil {
			log.Println("error unmarshalling location data:", err)
			return err
		}
	}

	fmt.Println(data)

	fmt.Println("Exploring", area+"...")
	if len(data.PokemonEncounters) == 0 {
		fmt.Println("No pokemon found...")
	}
	for _, v := range data.PokemonEncounters {
		fmt.Println("-", v.Pokemon.Name)
	}

	return nil
}
