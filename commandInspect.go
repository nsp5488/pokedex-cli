package main

import "fmt"

func commandInspect(c *Config, name string) error {
	pokemon, result := c.pokedex[name]
	if !result {
		fmt.Println("You have not caught that pokemon!")
		return nil
	}
	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n", pokemon.Name, pokemon.Height, pokemon.Weight)
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types: ")
	for _, t := range pokemon.Types {
		fmt.Printf("\t- %s\n", t.Type.Name)
	}

	return nil
}
