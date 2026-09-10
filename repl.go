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

func startRepl() error {
	commandList := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		command, exists := commandList[cleaned[0]]
		if !exists {
			fmt.Print("Unknown command\n")
		} else {
			command.callback()
		}
	}
}
