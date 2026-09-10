package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(input string) []string {
	words := []string{}
	lowered := strings.ToLower(input)
	words = strings.Fields(lowered)
	return words
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}
		command, exists := cfg.commands[cleaned[0]]
		if !exists {
			fmt.Print("Unknown command\n")
		} else {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}
