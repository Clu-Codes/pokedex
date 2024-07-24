package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/clu-codes/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	caughtPokemon map[string]pokeapi.Pokemon
	next          *string
	previous      *string
}

func startRepl(cfg *config) {
	fmt.Println("Initiated!...")
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		input := cleanText(reader.Text())
		if len(input) == 0 {
			continue
		}

		if len(input) == 1 {
			input = append(input, "")
		}
		cmd, city := input[0], input[1]

		cmdName, exists := getCommands()[cmd]
		if exists {
			err := cmdName.callback(cfg, city)
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

func cleanText(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Returns 20 locations within the Pokedex",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Returns the previous 20 locations. If no previous 20 exits, returns an error.",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore + {city}",
			description: "Returns the pokemon that exists in that city.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch + {pokemon_name}",
			description: "Throws a Pokeball at the desired pokemon, attempting to catch it.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect + {pokemon_name}",
			description: "Retrieves Pokemon from Pokedex and provides information about its types and abilities.",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Prints all caught Pokemon in your Pokedex",
			callback:    commandPokedex,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}
