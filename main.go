package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jming514/pokedex-go/internal/pokeapi"
)

var prompt string = "Pokedex > "

func main() {
	fmt.Println("Type 'help' to get all commands")
	input := bufio.NewScanner(os.Stdin)

	for {
		commandMap := map[string]func(){
			"exit":    func() { os.Exit(3) },
			"map":     pokeapi.Location.GetMap,
			"mapb":    pokeapi.Location.GetMapb,
			"explore": func() { fmt.Println("hi") },
		}

		fmt.Printf("%v", prompt)
		input.Scan()

		if strings.TrimSpace(input.Text()) == "help" {
			printCommands(commandMap)
			fmt.Println()
			continue
		}

		val, ok := commandMap[input.Text()]
		if ok {
			val()
		} else {
			fmt.Printf("Unrecongnized command: %v\n", input.Text())
		}
		fmt.Println()
	}
}

func printCommands(commands map[string]func()) {
	fmt.Println("Commands:")
	for k := range commands {
		fmt.Printf("%v\n", k)
	}
}
