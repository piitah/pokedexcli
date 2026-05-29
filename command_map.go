package main

import (
	"fmt"
)

func commandMapf(cfg *config, args []string) error {
	resp, err := cfg.pokeapiClient.GetLocationList(cfg.NextLocationURL)
	if err != nil {
		return fmt.Errorf("Something went wrong %w", err)
	}
	cfg.NextLocationURL = &resp.Next
	cfg.PreviousLocationURL = &resp.Previous

	for _, location := range resp.Results {
		fmt.Println("Request  ", location.Name)
	}

	return nil
}

func commandMapb(cfg *config, args []string) error {
	if cfg.PreviousLocationURL == nil || *cfg.PreviousLocationURL == "" {
		return fmt.Errorf("You are already on the first page")
	}
	resp, err := cfg.pokeapiClient.GetLocationList(cfg.PreviousLocationURL)
	if err != nil {
		return fmt.Errorf("Something went wrong %w", err)
	}
	cfg.NextLocationURL = &resp.Next
	cfg.PreviousLocationURL = &resp.Previous

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}

	return nil
}
