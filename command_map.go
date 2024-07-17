package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, s string) error {
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

func commandMapB(cfg *config, s string) error {
	if cfg.previous == nil {
		return errors.New("no previous route available")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.previous)
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
