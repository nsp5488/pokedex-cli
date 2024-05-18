package main

import (
	"time"

	"github.com/nsp5488/pokedexcli/internal/pokeapi"
)

func main() {
	config := Config{
		client: pokeapi.NewClient(time.Minute),
	}
	startRepl(config)
}
