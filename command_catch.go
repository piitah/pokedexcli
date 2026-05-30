package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Please provide pokemon name")
	}

	pokemon := args[0]

	resp, err := cfg.pokeapiClient.CatchPokemon(pokemon)
	if err != nil {
		return fmt.Errorf("something went wronggg %w ", err)
	}

	randomInt := rand.Intn(resp.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", resp.Name)
	if randomInt > 50 {
		fmt.Printf("%s escaped \n", resp.Name)
		return nil
	}
	fmt.Printf("%s was caught\n", resp.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	cfg.CaughtPokemon[resp.Name] = *resp
	return nil
}
