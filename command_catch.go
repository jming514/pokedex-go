package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(pokemonName ...string) error {
	fmt.Println("Catching pokemon: ", pokemonName)
	randomNumber := rand.Int()
	fmt.Print(randomNumber)
	return nil
}
