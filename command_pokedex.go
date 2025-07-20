package main

import (
	"fmt"
)

func commandPokedex(config *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, p := range config.caughtPokemon {
		fmt.Printf(" - %s\n", p.Name)
	}
	return nil
}
