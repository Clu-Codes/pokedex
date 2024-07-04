package main

import (
	"os"
)

func commandExit(c *config) error {
	os.Exit(0)
	return nil
}
