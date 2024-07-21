package pokedb

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Pokedex struct {
	pokedex map[string]pokemonEntry
	mu      *sync.Mutex
}

type pokemonEntry struct {
	caughtAt time.Time
	val      []byte
}

func (pdx *Pokedex) AddPokemon(key string, data []byte) error {
	pdx.mu.Lock()
	defer pdx.mu.Unlock()
	pdx.pokedex[key] = pokemonEntry{
		caughtAt: time.Now(),
		val:      data,
	}

	if _, ok := pdx.pokedex[key]; ok {
		fmt.Println("pokemon successfully added to pokedex")
	} else {
		return errors.New("failed to add pokemon to pokedex")
	}

	return nil
}

func (pdx *Pokedex) GetPokemonData(key string) ([]byte, bool) {
	pdx.mu.Lock()
	defer pdx.mu.Unlock()

	v, ok := pdx.pokedex[key]
	return v.val, ok
}

func LaunchPokedex() Pokedex {
	pdx := Pokedex{
		pokedex: make(map[string]pokemonEntry),
		mu:      &sync.Mutex{},
	}

	return pdx
}
