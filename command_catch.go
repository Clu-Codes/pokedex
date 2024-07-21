package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you need to provide the name of a pokemon")
	}

	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	exp := pokemon.BaseExperience
	digits := intLen(exp) * 10
	if digits > 90 {
		digits = 90
	}
	catch_chance := 100 - digits
	if catch_chance > rand.Intn(100) {
		fmt.Printf("%s was caught!\n", name)
		data, err := json.Marshal(pokemon)
		if err != nil {
			fmt.Printf("attempted to add to pokedex but incurred: %v\n", err)
		}
		cfg.pokedex.AddPokemon(name, data)

	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}

func intLen(i int) int {
	if i >= 1e18 {
		return 19
	}
	x, count := 10, 1
	for x <= i {
		x *= 10
		count++
	}
	return count
}
