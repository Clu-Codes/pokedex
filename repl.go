package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	cfg := config{
		next:     "",
		previous: "",
	}

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
			err := cmdName.callback(&cfg)
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

func printRes(c *config, resBody []byte) {
	// UnMarshalling the data allows mapping of resp data to pokeLoc struct, making it easy to access the resp data.
	loc := pokeLoc{}
	err := json.Unmarshal(resBody, &loc)
	if err != nil {
		fmt.Println(err)
	}

	// Update pointer value of config struct to next and previous locations, allowing pagination through api locations
	c.next = loc.Next
	c.previous = loc.Previous

	for _, v := range loc.Results {
		fmt.Println(v.Name)
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

type config struct {
	next     string
	previous any
}

type pokeLoc struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
