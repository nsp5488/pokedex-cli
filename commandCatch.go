package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(c *Config, name string) error {
	pokemon, err := c.client.CatchPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	roll := rand.Intn(max(400, pokemon.BaseExperience))
	if roll > pokemon.BaseExperience {
		c.pokedex[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", name)
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}
