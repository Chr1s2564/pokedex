package main

import (
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
)

type config struct {
	commands    map[string]cliCommand
	nextUrl     string
	previousUrl string
	client      *pokeapi.Client
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in Pokemon",
			callback:    displayMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the 20 previous location areas, if available",
			callback:    displayMapb,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for id, _ := range cfg.commands {
		fmt.Printf("%s: %s\n", cfg.commands[id].name, cfg.commands[id].description)
	}
	return nil
}
