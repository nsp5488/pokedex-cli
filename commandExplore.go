package main

import "fmt"

func commandExplore(c *Config, name string) error {
	pokemon, err := c.client.ExploreMap(name)
	if err != nil {
		return err
	}
	for _, mon := range pokemon {
		fmt.Printf(" - %s\n", mon)
	}
	return nil
}
