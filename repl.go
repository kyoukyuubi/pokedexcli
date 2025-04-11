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
		param := ""
		if len(words) > 1 {
			param = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, param)
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
	callback    func(*Config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
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
		"explore": {
			name:		 "explore <name of location>",
			description: "Get information about a location",
			callback: commandExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}