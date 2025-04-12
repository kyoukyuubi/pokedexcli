package main

import "fmt"

func commandInspect(config *Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please specify a caught pokemon name")
	}
	caughtPokemon := config.pokemonCaught
	name := args[0]

	pokemon, exists := caughtPokemon[name]
	if !exists {
		return fmt.Errorf("can't find pokemone with name: %s. Please catch it first using the 'catch' command", name)
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")

	for _, stats := range pokemon.Stats {
		fmt.Printf("-%s: %d\n", stats.Stat.Name, stats.BaseStat)
	}

	fmt.Println("Types:")
	
	for _, types := range pokemon.Types {
		fmt.Printf("-%s\n",types.Type.Name)
	}

	return nil
}