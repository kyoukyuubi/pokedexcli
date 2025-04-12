package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please specify a pokemon to catch, e.g., 'pikachu'")
	}

	pokemon, err := config.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return fmt.Errorf("error getting pokemon: %v", err)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	baseExperience := pokemon.BaseExperience

	catchChance := 100 - baseExperience/4
	if catchChance < 10 {
		catchChance = 10
	}

	randomNumber := rand.Intn(100) + 1

	if randomNumber <= catchChance {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		config.pokemonCaught[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}