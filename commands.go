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
	pokedex     map[string]pokeapi.CatchPokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
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
		"explore": {
			name:        "explore",
			description: "Explores the given location (takes a location argument)",
			callback:    displayExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch the pokemon given as argument",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects stats of a caught pokemon (takes a pokemon name as argument)",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Prints all the pokemons you have caught so far",
			callback:    commandPokedex,
		},
	}
}

func commandExit(cfg *config, arg string) error {
	if arg != "" {
		return fmt.Errorf("exit command doesn't take any argument")
	}
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, arg string) error {
	if arg != "" {
		return fmt.Errorf("help command doesn't take any argument")
	}
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for id, _ := range cfg.commands {
		fmt.Printf("%s: %s\n", cfg.commands[id].name, cfg.commands[id].description)
	}
	return nil
}
