package main

import (
	"time"
	"github.com/kyoukyuubi/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := &Config{
		pokeapiClient: pokeClient,
	}

	startRepl(config)
}
