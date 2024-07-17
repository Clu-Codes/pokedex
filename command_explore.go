package main

import (
	"fmt"
)

func commandExplore(cfg *config, city string) error {
	// fmt.Println(cfg.pokeapiClient.ListPokemon(city))
	fmt.Printf("Exploring %s...\n", city)
	pokemon, err := cfg.pokeapiClient.ListPokemon(city)
	if err != nil {
		fmt.Println("error found in commandExplore")
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemon.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
