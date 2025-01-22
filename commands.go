package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func SupportedCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the first 20 locations",
			callback:    commandMapBack,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

}

func commandExit(config *ConfigType) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *ConfigType) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range SupportedCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(config *ConfigType) error {
	if config.Next == "" {
		fmt.Println("you're on the last page")
		return nil
	}

	res, err := http.Get(config.Next)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	locationAreaResponse := LocationAreaResponse{}
	err = json.Unmarshal(body, &locationAreaResponse)

	if err != nil {
		return err
	}

	config.Next = ""
	config.Previous = ""
	if locationAreaResponse.Next != nil {
		config.Next = *locationAreaResponse.Next
	}

	if locationAreaResponse.Previous != nil {
		config.Previous = *locationAreaResponse.Previous
	}

	for _, locationArea := range locationAreaResponse.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}

func commandMapBack(config *ConfigType) error {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(config.Previous)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	locationAreaResponse := LocationAreaResponse{}
	err = json.Unmarshal(body, &locationAreaResponse)

	if err != nil {
		return err
	}

	fmt.Println(config)

	config.Next = ""
	config.Previous = ""
	if locationAreaResponse.Next != nil {
		config.Next = *locationAreaResponse.Next
	}

	if locationAreaResponse.Previous != nil {
		config.Previous = *locationAreaResponse.Previous
	}

	for _, locationArea := range locationAreaResponse.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}
