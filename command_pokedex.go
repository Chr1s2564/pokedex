package main

import (
	"fmt"
)

func commandPokedex(cfg *config, arg string) error {
	if arg != "" {
		return fmt.Errorf("pokedex command doesn't take any argument")
	}
	if len(cfg.pokedex) == 0 {
		return fmt.Errorf("your pokedex is empty, go catch some pokemons!")
	}
	fmt.Println("Your Pokedex")
	for _, pokemons := range cfg.pokedex {
		fmt.Printf(" - %s\n", pokemons.Name)
	}
	return nil
}
