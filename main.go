package main

import (
	"time"

	"github.com/nsp5488/pokedexcli/internal/pokeapi"
)

func main() {
	config := Config{
		client:  pokeapi.NewClient(time.Minute),
		pokedex: make(map[string]pokeapi.Pokemon),
	}
	startRepl(config)
}
