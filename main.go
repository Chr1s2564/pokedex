package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

func main() {
	cfg := &config{
		commands:    getCommands(),
		nextUrl:     "https://pokeapi.co/api/v2/location-area/",
		previousUrl: "",
		client:      pokeapi.NewClient(5 * time.Second),
	}
	startRepl(cfg)
}
