package main

import (
	// "encoding/json"
	"fmt"
	"io"
	"net/http"
	// "github.com/mtslzr/pokeapi-go"
)

func commandMap(c *config) error {
	// Use config struct to paginate through location api, starting with default url
	var url string
	if c.previous == "" {
		url = "https://pokeapi.co/api/v2/location"
	} else {
		url = c.next
	}

	// TODO: Look into use of pokeapi-go lib here instead of manual call
	res, err := http.Get(url)

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

	printRes(c, body)

	fmt.Println()

	return nil
}
