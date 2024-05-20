package main

import "os"

func commandExit(c *Config, name string) error {
	os.Exit(0)
	return nil
}
