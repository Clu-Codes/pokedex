package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		input := cleanText(reader.Text())
		if len(input) == 0 {
			continue
		}

		cmd := input[0]

		cmdName, exists := getCommands()[cmd]
		if exists {
			err := cmdName.callback()
			if err != nil { // redundant given schema for cliCommand struct, but keeping just in case.
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
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
