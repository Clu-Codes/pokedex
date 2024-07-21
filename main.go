package main

import (
	"time"

	"github.com/clu-codes/pokedex/internal/pokeapi"
	"github.com/clu-codes/pokedex/internal/pokedb"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	pokedb := pokedb.LaunchPokedex()
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       pokedb,
	}
	startRepl(cfg)
}
