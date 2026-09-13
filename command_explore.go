package main

import (
	"fmt"
)

func displayExplore(cfg *config, location string) error {
	if location == "" {
		return fmt.Errorf("explore command should have a location argument")
	}
	fmt.Printf("Exploring %s\n", location)
	pokemons, err := cfg.client.GetPokemons(location)
	if err != nil {
		return fmt.Errorf("error retrieving pokemons: %v", err)
	}
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemons.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
