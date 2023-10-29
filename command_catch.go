package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(pokemonName ...string) error {
	if len(pokemonName) == 0 || pokemonName[0] == "" {
		fmt.Println("Please enter a pokemon name")
		return errors.New("no pokemon name provided")
	}

	fmt.Println("Catching pokemon: ", pokemonName)
	randomNumber := rand.Int()
	fmt.Print(randomNumber)
	return nil
}
