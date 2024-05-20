package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nsp5488/pokedexcli/internal/pokeapi"
)

func startRepl(config Config) {
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	var argument string
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

		if len(input) > 1 {
			argument = input[1]
		}
		err := command.callback(&config, argument)
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
	callback    func(*Config, string) error
}

type Config struct {
	next     *string
	previous *string
	client   *pokeapi.Client
}

func getCommands() map[string]cliCommand {
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
		description: "Displays the previous 20 location areas in the Pokemon world",
		callback:    commandMapB,
	}
	commands["explore"] = cliCommand{
		name:        "explore",
		description: "Explores the passed location and lists the Pokemon available in that zone",
		callback:    commandExplore,
	}

	return commands
}
