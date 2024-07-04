package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandMapB(c *config) error {
	var url interface{}
	err := errors.New("No previous route available.")
	if c.previous == nil {
		fmt.Println(err)
		return nil
	}

	url = c.previous
	res, err := http.Get(url.(string))

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}

	printRes(c, body)

	fmt.Println()

	return nil
}
