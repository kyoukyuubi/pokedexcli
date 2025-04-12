package main

import "fmt"

func commandPokedex(config *Config, args ...string) error {
	caughtPokemon := config.pokemonCaught
	if len(caughtPokemon) < 1 {
		return fmt.Errorf("you haven't caught any Pokemon yet")
	}

	fmt.Println("Your Pokedex: ")
	for _, pokemon := range caughtPokemon {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}