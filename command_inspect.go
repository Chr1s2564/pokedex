package main

import (
	"fmt"
)

func commandInspect(cfg *config, arg string) error {
	if arg == "" {
		return fmt.Errorf("inspect command takes a pokemon's name as argument")
	}
	pokemonName, ok := cfg.pokedex[arg]
	if !ok {
		return fmt.Errorf("you have not caught %s", arg)
	}
	fmt.Printf("Name: %s\n", pokemonName.Name)
	fmt.Printf("Height: %d\n", pokemonName.Height)
	fmt.Printf("Weight: %d\n", pokemonName.Weight)
	fmt.Println("Stats:")
	for _, stats := range pokemonName.Stats {
		fmt.Printf(" - %s: %d\n", stats.Stat.Name, stats.BaseStat)
	}
	fmt.Println("Types: ")
	for _, types := range pokemonName.Types {
		fmt.Printf(" - %s\n", types.PokeTypes.Name)
	}

	return nil
}
