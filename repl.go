package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/piitah/pokedexcli/internal/pokeapi"
	pokecache "github.com/piitah/pokedexcli/internal/pokecache"
)

type config struct {
	cache               pokecache.Cache
	pokeapiClient       *pokeapi.Client
	NextLocationURL     *string
	PreviousLocationURL *string
}
type clicommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		word := clearInput(scanner.Text())

		if word == nil {
			continue
		}
		if cmd, exist := getCommands()[word[0]]; exist {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		fmt.Printf("Unknown command: %v\n", strings.TrimRight(word[0], "\r\n"))

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}
}

func getCommands() map[string]clicommand {
	commands := map[string]clicommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex CLI",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapb,
		},
	}
	return commands
}
func clearInput(s string) []string {
	result := strings.Fields(strings.ToLower(s))
	return result
}
