package main

import (
	"time"

	pokeapi "github.com/piitah/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       *pokeapi.Client
	NextLocationURL     *string
	PreviousLocationURL *string
	CaughtPokemon       map[string]pokeapi.PokemonResponse
}

func main() {
	pokeClient := pokeapi.NewClient(10 * time.Second)
	cfg := config{
		CaughtPokemon: make(map[string]pokeapi.PokemonResponse),
		pokeapiClient: pokeClient,
	}
	startRepl(&cfg)
}
