package main

import (
	"fmt"
	"strings"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Please provide explore location name")
	}

	areaName := strings.Join(args, "-")

	resp, error := cfg.pokeapiClient.GetLocationArea(areaName)
	if error != nil {
		return fmt.Errorf("Something went wrong %w", error)
	}

	fmt.Println("Exploring pastoria-city-area...")
	fmt.Println("Found Pokemon:")
	for _, location := range resp.PokemonEncounters {
		fmt.Println("- " + location.Pokemon.Name)
	}

	return nil
}
