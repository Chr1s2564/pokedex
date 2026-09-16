package main

import (
	"fmt"
)

func commandCatch(cfg *config, arg string) error {
	if arg == "" {
		return fmt.Errorf("catch command requires an argument (pokemon name)")
	}

	hasCatched, foundPokemon, err := cfg.client.CatchPokemon(arg)
	if err != nil {
		return fmt.Errorf("error finding pokemon: %v", err)
	}
	if !hasCatched {
		return fmt.Errorf("%s escaped!", arg)
	}
	cfg.pokedex[arg] = foundPokemon
	fmt.Printf("%s was caught!\n", arg)
	return nil
}
