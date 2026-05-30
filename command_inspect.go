package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("Please provide inspect name")
	}

	pokemon := args[0]

	result, exist := cfg.CaughtPokemon[pokemon]
	if exist {
		fmt.Printf("Name: %s\n", result.Name)
		fmt.Printf("Height: %d\n", result.Height)
		fmt.Printf("Weight: %d\n", result.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range result.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, typeInfo := range result.Types {
			fmt.Printf("  - %s\n", typeInfo.Type.Name)
		}
		return nil
	}
	fmt.Println("you have not caught that pokemon")
	return nil
}
