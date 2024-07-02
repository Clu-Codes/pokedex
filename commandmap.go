package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	// "github.com/mtslzr/pokeapi-go"
)

func commandMap() error {
	// May be cleaner to use pokeapi-go lib here instead of manual call
	res, err := http.Get("https://pokeapi.co/api/v2/location")

	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		fmt.Println(err)
	}

	// UnMarshalling the data allows mapping of resp data to pokeLoc struct, making it easy to access the resp data.
	// TODO: Loop through the Results slice to print the names of each city
	// TODO: Understand how to use instanced pokeLoc data for commandMapB
	loc := pokeLoc{}
	er := json.Unmarshal(body, &loc)
	if er != nil {
		fmt.Println(er)
	}

	return nil
}

type pokeLoc struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
