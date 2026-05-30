package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {
	fmt.Println("Your Pokedex:")
	for key, _ := range cfg.CaughtPokemon {
		fmt.Printf("  - %s\n", key)
	}
	return nil
}
