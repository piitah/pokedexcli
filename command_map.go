package main

import (
	"encoding/json"
	"fmt"

	pokeapi "github.com/piitah/pokedexcli/internal/pokeapi"
)

func commandMapf(cfg *config) error {
	args := cfg.NextLocationURL

	if args != nil {
		cachedata, ok := cfg.cache.Get(*cfg.NextLocationURL)
		if ok {
			resp := pokeapi.LocationListResponse{}
			if err := json.Unmarshal(cachedata, &resp); err != nil {
				return fmt.Errorf("Something went wrong %w", err)
			}
			cfg.NextLocationURL = &resp.Next
			cfg.PreviousLocationURL = &resp.Previous

			for _, location := range resp.Results {
				fmt.Println(location.Name)
			}
			return nil
		}
	}

	resp, err := cfg.pokeapiClient.GetLocationList(cfg.NextLocationURL)
	if err != nil {
		return fmt.Errorf("Something went wrong %w", err)
	}
	cfg.NextLocationURL = &resp.Next
	cfg.PreviousLocationURL = &resp.Previous

	marshaled, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("Something went wrong %w", err)
	}
	cfg.cache.Add(*cfg.NextLocationURL, marshaled)

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.PreviousLocationURL == nil {
		return fmt.Errorf("You are already on the first page")
	}

	cacheData, ok := cfg.cache.Get(*cfg.PreviousLocationURL)
	if ok {
		var resp pokeapi.LocationListResponse

		if err := json.Unmarshal(cacheData, &resp); err != nil {
			return fmt.Errorf("Something went wrong %w", err)
		}

		for _, location := range resp.Results {
			fmt.Println(location.Name)
		}
		return nil
	}

	resp, err := cfg.pokeapiClient.GetLocationList(cfg.PreviousLocationURL)
	if err != nil {
		return fmt.Errorf("Something went wrong %w", err)
	}
	cfg.NextLocationURL = &resp.Next
	cfg.PreviousLocationURL = &resp.Previous

	marshaled, err := json.Marshal(resp)
	if err != nil {
		return fmt.Errorf("Something went wrong %w", err)
	}

	cfg.cache.Add(*cfg.PreviousLocationURL, marshaled)
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}

	return nil
}
