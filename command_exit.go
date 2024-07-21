package main

import (
	"os"
)

func commandExit(c *config, args ...string) error {
	os.Exit(0)
	return nil
}
