package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {

	/*
		if cache, exists := cfg.cache.GetCache(*cfg.next); exists {
			// TODO: Unmarhsal []byte from cache to PokeLoc{}

			// cfg.next = cache.Next
			// cfg.previous = cache.Previous

			// for _, loc := range cache.Results {
			// 	fmt.Println(loc.Name)
			// }
		}
	*/

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.next)
	if err != nil {
		return err
	}

	cfg.next = locationsResp.Next
	cfg.previous = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapB(cfg *config) error {
	if cfg.previous == nil {
		return errors.New("no previous route available")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.previous)
	if err != nil {
		return err
	}

	cfg.next = locationResp.Next
	cfg.previous = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
