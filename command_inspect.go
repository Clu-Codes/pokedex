package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you need to provide the name of a pokemon")
	}

	name := args[0]
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		fmt.Println("pokemon not in pokedex. catch pokemon to add to pokedex")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")

	for _, stats := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stats.Stat.Name, stats.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, types := range pokemon.Types {
		fmt.Printf("  -%s\n", types.Type.Name)
	}

	return nil
}
