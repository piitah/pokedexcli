package main

import "fmt"

func commandHelp(cfg *config, args []string) error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")
	for _, cmd := range getCommands() {
		fmt.Printf("%s  - %s\n", cmd.name, cmd.description)
	}
	fmt.Print("  \n")
	return nil
}
