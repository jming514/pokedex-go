package main

import (
	"bufio"
	"fmt"
	"github.com/jming514/pokedex-go/internal/pokeapi"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "exit cli",
			callback:    commandExitCli,
		},
		"map": {
			name:        "map",
			description: "move forward maps",
			callback:    nil,
		},
		"mapb": {
			name:        "mapb",
			description: "move backward maps",
			callback:    nil,
		},
		"explore": {
			name:        "explore",
			description: "does nothing right now",
			callback:    nil,
		},
		"help": {
			name:        "help",
			description: "list all the commands",
			callback:    commandHelp,
		},
	}
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			_ = command.callback()
		} else {
			fmt.Println("Unrecognized command")
			fmt.Println()
		}

		fmt.Println()
	}
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowered := strings.ToLower(trimmed)
	output := strings.Fields(lowered)
	return output
}
