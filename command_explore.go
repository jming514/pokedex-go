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

	exploreUrl := locationUrl + "-area/" + area
	fmt.Println(exploreUrl)
	res, err := http.Get(exploreUrl)
	if err != nil {
		log.Println("error fetching location data:", err)
		return err
	}

	if strings.Contains(res.Status, "404") {
		return errors.New(res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("error reading res.body:", err)
		return err
	}
	defer res.Body.Close()

	data := exploreResponse{}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("error unmarshalling location data:", err)
		return err
	}

	fmt.Println("Exploring", area+"...")
	for _, v := range data.PokemonEncounters {
		fmt.Println("-", v.Pokemon.Name)
	}

	cfg.pokeapiClient.C.Add(exploreUrl, body)

	return nil
}
