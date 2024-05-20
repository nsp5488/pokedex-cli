package main

import (
	"errors"
	"fmt"
)

func commandMap(c *Config, name string) error {
	result, err := c.client.GetMapFromPokeApi(c.next)

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

func commandMapB(c *Config, name string) error {
	result, err := c.client.GetMapFromPokeApi(c.previous)
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
