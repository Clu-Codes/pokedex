package main

import (
	"time"

	"github.com/clu-codes/pokedex/internal/pokeapi"
	"github.com/clu-codes/pokedex/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(10 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		cache:         pokeCache,
	}
	startRepl(cfg)
}
