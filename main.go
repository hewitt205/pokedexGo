package main

import (
	"time"

	"github.com/hewitt205/Pokedex/internal/pokeapi"
)


func main() {
	pokeClient := pokeapi.newClient(5 * time.Second)
	config := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(config)
}
