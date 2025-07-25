package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hewitt205/pokedexGo/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon map[string]pokeapi.Pokemon
}


func startRepl(config *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name: "map",
			description: "Lists the names of world locations, 20 at a time",
			callback: commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Lists the names of the previous 20 world locations",
			callback: commandMapb,
		},
		"inspect": {
			name: "inspect <pokemon_name>",
			description: "View details about a caught Pokemon",
			callback: commandInspect,
		},
		"explore": {
			name: "explore <location_name>",
			description: "Explore a location",
			callback: commandExplore,
		},
		"catch": {
			name: "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback: commandCatch,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"pokedex": {
			name: "pokedex",
			description: "See all the pokemon you've caught",
			callback: commandPokedex,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
	}
}
