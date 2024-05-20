package main

import "fmt"

func commandPokedex(c *Config, _ string) error {
	fmt.Printf("You have caught %d pokemon!\n", len(c.pokedex))
	for _, value := range c.pokedex {
		fmt.Printf("\t -%s\n", value.Name)
	}
	return nil
}
