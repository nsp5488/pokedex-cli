package main

import "os"

func commandExit(c *Config) error {
	os.Exit(0)
	return nil
}
