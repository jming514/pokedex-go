package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jming514/pokedex-go/internal/pokeapi"
)

type callbackType func(args ...string) error

type cliCommand struct {
	callback    callbackType
	name        string
	description string
}

type config struct {
	pokeapiClient   *pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": {
			name:        "catch",
			description: "catch pokemon",
			callback:    commandCatch,
		},
		"exit": {
			name:        "exit",
			description: "exit cli",
			callback:    commandExitCli,
		},
		"map": {
			name:        "map",
			description: "move forward maps",
			callback:    cfg.commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "move backward maps",
			callback:    cfg.commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "does nothing right now",
			callback:    cfg.commandExplore,
		},
		"help": {
			name:        "help",
			description: "list all the commands",
			callback:    commandHelp,
		},
		"dc": {
			name:        "cd",
			description: "debug cache",
			callback:    cfg.commandDebugCache,
		},
	}
}

var cfg = config{}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	client := pokeapi.NewClient(5 * time.Minute)
	cfg.pokeapiClient = &client
	for {
		fmt.Printf("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		commandName := words[0]
		commandArgs := words[1:]

		command, exists := getCommands()[commandName]
		if exists {
			_ = command.callback(commandArgs...)
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

func printResults(input []struct {
	Name string
	URL  string
},
) {
	for _, v := range input {
		fmt.Printf("%s\n", v.Name)
	}
}
