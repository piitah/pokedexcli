package main

import (
	"time"

	pokeapi "github.com/piitah/pokedexcli/internal/pokeapi"
	pokecache "github.com/piitah/pokedexcli/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(10 * time.Second)
	cache := pokecache.NewCache(10 * time.Second)
	cfg := config{
		cache:         cache,
		pokeapiClient: pokeClient,
	}
	startRepl(&cfg)
}
