package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var locationUrl = "https://pokeapi.co/api/v2/location/"

func GetMap() (PokedexLocation, error) {
	apiUrl := locationUrl

	res, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		log.Printf("error sending http request: %v\n", err)
		return PokedexLocation{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("error reading body: %v\n", err)
		return PokedexLocation{}, err
	}

	err = res.Body.Close()
	if err != nil {
		log.Printf("error closing body: %v\n", err)
		return PokedexLocation{}, err
	}

	locationResponse := PokedexLocation{}
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		log.Printf("cannot unmarshal: %v\n", err)
		return PokedexLocation{}, err
	}

	return locationResponse, nil
}
