package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jming514/pokedex-go/internal/pokeapi"
)

var prompt string = "Pokedex > "

func main() {
	input := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("%v", prompt)
		input.Scan()

		switch input.Text() {
		case "exit":
			os.Exit(3)

		case "map":
			pokeapi.GetMap()

		case "mapb":
			pokeapi.GetMapb()

		case "help":
			fmt.Println("command\tdescription\nhelp\tprint out all commands\nexit\tclose the program")

		default:
			fmt.Printf("Unrecongnized command: %v\n", input.Text())
		}

		fmt.Println()
	}
}
