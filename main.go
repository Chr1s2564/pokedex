package main

//import "fmt"

func main() {
	cfg := &config{
		commands: getCommands(),
	}
	startRepl(cfg)
}
