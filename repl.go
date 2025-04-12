package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kyoukyuubi/pokedexcli/internal/pokeapi"
)

type Config struct {
	pokeapiClient pokeapi.Client
    nextLocationsURL *string
    prevLocationsURL *string
	pokemonCaught map[string]pokeapi.Pokemon
}

func startRepl(config *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"catch": {
			name:		 "catch <name of pokemon>",
			description: "Attempt to catch a Pokemon",
			callback: commandCatch,
		},
		"explore": {
			name:		 "explore <name of location>",
			description: "Get information about a location",
			callback: commandExplore,
		},
		"inspect": {
			name:		 "inspect <name of location>",
			description: "Get information about a caught pokemon",
			callback: commandInspect,
		},
		"map": {
			name:		 "map",
			description: "Get the next page of locations",
			callback: commandMap,
		},
		"mapb": {
			name:		 "mapb",
			description: "Get the previous page of locations",
			callback: commandMapBack,
		},
		"pokedex": {
			name:		 "inspect",
			description: "Get a list of the Pokemon you've caught",
			callback: commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}