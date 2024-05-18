package main

import (
	"errors"
	"fmt"
	"internal/pokeapi"
)

func commandMap(c *Config) error {
	result, err := pokeapi.GetMapFromPokeApi(c.next)

	if err != nil {
		return errors.New("error while fetching")
	}

	c.next = result.Next
	c.previous = result.Previous

	for _, location := range result.Locations {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}

func commandMapB(c *Config) error {
	result, err := pokeapi.GetMapFromPokeApi(c.previous)
	if err != nil {
		return errors.New("error while fetching")
	}

	c.next = result.Next
	c.previous = result.Previous

	for _, location := range result.Locations {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}
