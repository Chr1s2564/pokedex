package main

//import "fmt"

func main() {
	cfg := &config{
		commands:    getCommands(),
		nextUrl:     "https://pokeapi.co/api/v2/location-area/",
		previousUrl: "",
	}
	startRepl(cfg)
}
