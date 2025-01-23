package main

import (
	"fmt"
	"os"

	"github.com/Jelle-S/pokedexcli/internal/api"

	"github.com/Jelle-S/pokedexcli/models"
)

func supportedCommands() map[string]models.CliCommand {
	return map[string]models.CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Display the next 20 locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Display the first 20 locations",
			Callback:    commandMapBack,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}

}

func commandExit(config *models.ConfigType) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *models.ConfigType) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range supportedCommands() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}

func commandMap(config *models.ConfigType) error {
	if config.Next == "" {
		fmt.Println("you're on the last page")
		return nil
	}

	locationAreaResponse, err := api.GetAndUnmarshal[models.LocationAreaResponse](config.Next, config.Cache)

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

func commandMapBack(config *models.ConfigType) error {
	if config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locationAreaResponse, err := api.GetAndUnmarshal[models.LocationAreaResponse](config.Previous, config.Cache)

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
