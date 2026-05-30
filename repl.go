package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type clicommand struct {
	name        string
	description string
	callback    func(cfg *config, args []string) error
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
		args := word[1:]
		if cmd, exist := getCommands()[word[0]]; exist {
			err := cmd.callback(cfg, args)
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
			description: "Get the previous page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the next page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "explore",
			description: "Explore location area",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "explore",
			description: "Explore location area",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "explore",
			description: "Explore location area",
			callback:    commandPokedex,
		},
	}
	return commands
}
func clearInput(s string) []string {
	result := strings.Fields(strings.ToLower(s))
	return result
}
