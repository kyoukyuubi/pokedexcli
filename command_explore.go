package main

import (
	"fmt"
	"strings"
)

func commandExplore(config *Config, param string) error {
	if param == "" {
		return fmt.Errorf("please specify a location to explore, e.g., 'explore pastoria-city-area'")
	}
	locationsResp, err := config.pokeapiClient.Location(param)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return fmt.Errorf("location '%s' not found; please check the name and try again", param)
		}
		return fmt.Errorf("unexpected error fetching location: %v", err)
	}

	fmt.Printf("Exploring %s...\nFound Pokémon:\n", locationsResp.Name)
	for _, pokemon := range locationsResp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}