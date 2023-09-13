package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
	Next     string
	Previous string
}

func (c config) getMap() {
	res, err := http.Get("https://pokeapi.co/api/v2/location/")
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
	res.Body.Close()

	location = config{
		Next:     t.Next,
		Previous: t.Previous,
	}

	fmt.Println("location", location)
}

var prompt string = "Pokedex > "

var location config = config{
	Next:     "",
	Previous: "",
}

func main() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("%v", prompt)
		input.Scan()

		switch input.Text() {
		case "exit":
			os.Exit(3)

		case "map":
			location.getMap()

		case "help":
			fmt.Println("command\tdescription\nhelp\tprint out all commands\nexit\tclose the program")

		default:
			fmt.Printf("Unrecongnized command: %v\n", input.Text())
		}

		fmt.Println()
	}
}
