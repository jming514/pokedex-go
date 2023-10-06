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

var location = config{
	Next:     nil,
	Previous: nil,
}

func random() {
	fmt.Println("This is something random")
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

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	t := pokedexLocation{}

	err = json.Unmarshal(body, &t)
	if err != nil {
		log.Fatal(err)
	}
	err = res.Body.Close()
	if err != nil {
		return
	}

	location.Next = &t.Next
	location.Previous = &t.Previous

	for _, v := range t.Results {
		fmt.Println(v.Name)
	}
}

func GetMapb() {
	locationUrl := "https://pokeapi.co/api/v2/location/"
	if location.Previous != nil {
		//if *location.Previous != "" {
		locationUrl = *location.Previous
		//}
	}

	res, err := http.Get(locationUrl)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	t := pokedexLocation{}

	err = json.Unmarshal(body, &t)
	if err != nil {
		log.Fatal(err)
	}
	err = res.Body.Close()
	if err != nil {
		return
	}

	location.Next = &t.Next
	location.Previous = &t.Previous

	for _, v := range t.Results {
		fmt.Println(v.Name)
	}
}
