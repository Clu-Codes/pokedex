package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("There are no pokemon in the pokedex. Go catch some!")
	}
	for name := range cfg.caughtPokemon {
		fmt.Printf(" -%s\n", name)
	}

	return nil
}
