package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	commands := SupportedCommands()
	scanner := bufio.NewScanner(os.Stdin)
	config := ConfigType{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())[0]
		command, ok := commands[input]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(&config)
		if err != nil {
			panic(err)
		}
	}
}

func cleanInput(text string) []string {
	result := []string{}
	for _, s := range strings.Fields(text) {
		result = append(result, strings.ToLower(strings.TrimSpace(s)))
	}
	return result
}
