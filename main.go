package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	commands := supportedCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())[0]
		command, ok := commands[input]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		command.callback()
	}
}

func cleanInput(text string) []string {
	result := []string{}
	for _, s := range strings.Fields(text) {
		result = append(result, strings.ToLower(strings.TrimSpace(s)))
	}
	return result
}

func supportedCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range supportedCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
