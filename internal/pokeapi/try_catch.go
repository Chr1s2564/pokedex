package pokeapi

import (
	"fmt"
	"math/rand"
)

type UserPokedex struct {
	CaughtPokemon map[string]CatchPokemon
}

func (c *Client) CatchPokemon(pokemon string) (bool, CatchPokemon, error) {
	foundPokemon, err := c.FindPokemon(pokemon)
	if err != nil {
		return false, CatchPokemon{}, err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	chance := 10000 / (foundPokemon.BaseExperience + 100)
	roll := rand.Intn(100)
	if roll < chance {
		return true, foundPokemon, nil
	}
	return false, CatchPokemon{}, nil
}
