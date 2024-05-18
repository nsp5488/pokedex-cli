package main

import (
	"bufio"
	"errors"
	"fmt"
	"internal/pokeapi"
	"os"
	"strings"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommandMap()
	config := buildConfig()

	for {
		fmt.Print("pokedex > ")
		reader.Scan()
		input := cleanInput(reader.Text())
		if len(input) == 0 {
			continue
		}

		command, ok := commands[input[0]]

		if !ok {
			continue
		}
		err := command.callback(&config)
		if err != nil {
			break
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)

}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	next     *string
	previous *string
}

func buildConfig() Config {
	next := "https://pokeapi.co/api/v2/location-area/"
	previous := "test"
	return Config{
		next:     &next,
		previous: &previous,
	}
}

func getCommandMap() map[string]cliCommand {
	commands := make(map[string]cliCommand)

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}

	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays the names of 20 location areas in the Pokemon world",
		callback:    commandMap,
	}
	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "displays the previous 20 location areas in the Pokemon world",
		callback:    commandMapB,
	}

	return commands
}

func commandExit(c *Config) error {
	os.Exit(0)
	return nil
}

func commandHelp(c *Config) error {
	commands := getCommandMap()
	for value := range commands {
		command := commands[value]
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(c *Config) error {
	if c.next == nil {
		return errors.New("no next locations to fetch")
	}
	result, err := pokeapi.GetMapFromPokeApi(*c.next)
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
	if c.previous == nil {
		return errors.New("no previous locations to fetch")
	}
	result, err := pokeapi.GetMapFromPokeApi(*c.previous)
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
