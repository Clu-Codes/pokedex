package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/clu-codes/pokedex/internal/pokeapi"
)

// const (
// 	pokemonDetail = `
// 	Name: %s
// 	Height: %d
// 	Weight: %d
// 	Stats:
// 		-hp: %d
// 		-attack: %d
// 		-defense: %d
// 		-special-attack: %d
// 		-special-defense: %d
// 		-speed: %d
// 	Types:
// 		- %s
// 		- %s
// 	`
// )

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you need to provide the name of a pokemon")
	}

	name := args[0]
	pokemon, ok := cfg.pokedex.GetPokemonData(name)
	if !ok {
		fmt.Println("pokemon not in pokedex. catch pokemon to add to pokedex")
		return nil
	}
	data := pokeapi.Pokemon{}
	err := json.Unmarshal(pokemon, &data)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\n", data.Name)
	fmt.Printf("Height: %d\n", data.Height)
	fmt.Printf("Weight: %d\n", data.Weight)
	fmt.Printf("Stats:\n")

	for _, stats := range data.Stats {
		// statsObj := make(map[string]int)
		// key := stats.Stat.Name
		// statsObj[key] = stats.BaseStat
		fmt.Printf("  -%s: %d\n", stats.Stat.Name, stats.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, types := range data.Types {
		fmt.Printf("  -%s\n", types.Type.Name)
	}

	// TODO: retrieve pokemon data from JSON and format pokemon data into string
	// May need to loop through data.Stats to get data correctly
	// fmt.Printf(pokemonDetail, data.Name, data.Height, data.Weight, statsObj["hp"], statsObj["attack"], statsObj["defense"], statsObj["special-attack"], statsObj["special-attack"], statsObj["speed"], )

	return nil
}
