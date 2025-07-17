package main

import (
	"time"

	"github.com/hewitt205/pokedexGo/internal/pokeapi"
)


func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(config)
}
