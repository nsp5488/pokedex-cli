package main

import "fmt"

func commandHelp(c *Config) error {
	commands := getCommands()
	for value := range commands {
		command := commands[value]
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
