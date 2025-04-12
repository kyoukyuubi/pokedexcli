package main

import (
	"fmt"
)

func commandExplore(config *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please specify a location to explore, e.g., 'explore pastoria-city-area'")
	}
	name := args[0]
	location, err := config.pokeapiClient.Location(name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}