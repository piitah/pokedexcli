package main

import (
	"time"

	pokeapi "github.com/piitah/pokedexcli/internal/pokeapi"
	pokecache "github.com/piitah/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cache := pokecache.NewCache(5 * time.Second)
	cfg := config{
		cache:         cache,
		pokeapiClient: pokeClient,
	}
	startRepl(&cfg)
}
