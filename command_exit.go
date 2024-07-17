package main

import (
	"os"
)

func commandExit(c *config, s string) error {
	os.Exit(0)
	return nil
}
