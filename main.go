package main

import (
	"time"

	"github.com/kyoukyuubi/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	config := &Config{
		pokeapiClient: pokeClient,
	}
	
	startRepl(config)
}
