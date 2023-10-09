package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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

var location = &config{
	Next:     nil,
	Previous: nil,
}

func random() {
	fmt.Println("This is something random")
}

func (cfg *config) setNextPrev(res *http.Response) error {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	pl := pokedexLocation{}

	err = json.Unmarshal(body, &pl)
	if err != nil {
		log.Fatal(err)
	}
	err = res.Body.Close()
	if err != nil {
		return err
	}

	cfg.Next = &pl.Next
	cfg.Previous = &pl.Previous

	for _, v := range pl.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func GetMap() {
	locationUrl := "https://pokeapi.co/api/v2/location/"
	if location.Next != nil {
		locationUrl = *location.Next
	}

	res, err := http.Get(locationUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = location.setNextPrev(res)
	if err != nil {
		log.Println(err)
	}
}

func GetMapb() {
	locationUrl := "https://pokeapi.co/api/v2/location/"
	if location.Previous != nil {
		locationUrl = *location.Previous
	}

	res, err := http.Get(locationUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = location.setNextPrev(res)
	if err != nil {
		log.Println(err)
	}
}
