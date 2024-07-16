package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/clu-codes/pokedex/internal/pokeapi"
)

func commandMap(cfg *config) error {
	if cfg.next == nil {
		fmt.Println("cache miss; key does not exist in cache")
	} else {
		if cache, exists := cfg.cache.GetCache(*cfg.next); exists {
			// TODO: Unmarhsal []byte from cache to PokeLoc{}
			lData := pokeapi.PokeLoc{}
			err := json.Unmarshal(cache, &lData)
			if err != nil {
				fmt.Printf("error unmarshalling cache data: %v", err)
			}

			fmt.Println("pulled from cache")

			cfg.next = lData.Next
			cfg.previous = lData.Previous

			for _, loc := range lData.Results {
				fmt.Println(loc.Name)
			}

			return nil
		} else {
			fmt.Println("no cfg.next url exists in cache")
		}
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.next)
	if err != nil {
		return err
	}

	mData, err := json.Marshal(locationsResp)
	if err != nil {
		fmt.Printf("unable to cache response due to %v", err)
	}

	url_key := pokeapi.BaseURL + "/location-area"
	if cfg.next != nil {
		url_key = *cfg.next
	}

	cfg.cache.AddCache(url_key, mData)

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

	if cache, exists := cfg.cache.GetCache(*cfg.previous); exists {
		// TODO: Unmarhsal []byte from cache to PokeLoc{}
		lData := pokeapi.PokeLoc{}
		err := json.Unmarshal(cache, &lData)
		if err != nil {
			fmt.Printf("error unmarshalling cache data: %v", err)
		}

		fmt.Println("pulled from cache")

		cfg.next = lData.Next
		cfg.previous = lData.Previous

		for _, loc := range lData.Results {
			fmt.Println(loc.Name)
		}

		return nil

	} else {
		fmt.Println("no cfg.previous url exists in cache. querying results...")
		locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.previous)
		if err != nil {
			return err
		}

		mData, err := json.Marshal(locationsResp)
		if err != nil {
			fmt.Printf("unable to cache response due to %v", err)
		}
		cfg.cache.AddCache(*cfg.previous, mData)

		cfg.next = locationsResp.Next
		cfg.previous = locationsResp.Previous

		for _, loc := range locationsResp.Results {
			fmt.Println(loc.Name)
		}

		return nil

	}
}
