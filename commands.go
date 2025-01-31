package main

import (
	"fmt"
	"math/rand"
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
		"explore": {
			Name:        "explore",
			Description: "Explore an area",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a pokemon (or try to)",
			Callback:    commandCatch,
		},
	}

}

func commandExit(config *models.ConfigType, arguments []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *models.ConfigType, arguments []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range supportedCommands() {
		fmt.Printf("%s: %s\n", command.Name, command.Description)
	}
	return nil
}

func commandMap(config *models.ConfigType, arguments []string) error {
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

func commandMapBack(config *models.ConfigType, arguments []string) error {
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

func commandExplore(config *models.ConfigType, arguments []string) error {
	fmt.Println("Exploring " + arguments[0] + "...")

	locationArea, err := api.GetAndUnmarshal[models.LocationArea]("https://pokeapi.co/api/v2/location-area/"+arguments[0], config.Cache)

	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, encounter := range locationArea.PokemonEncounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(config *models.ConfigType, arguments []string) error {
	fmt.Println("Throwing a Pokeball at " + arguments[0] + "...")

	Pokemon, err := api.GetAndUnmarshal[models.Pokemon]("https://pokeapi.co/api/v2/pokemon/"+arguments[0], config.Cache)

	if err != nil {
		return err
	}

	if rand.Intn(Pokemon.BaseExp) > (Pokemon.BaseExp - 30) {
		config.Pokedex[Pokemon.Name] = Pokemon
		fmt.Println(Pokemon.Name + " was caught!")
		return nil
	}

	fmt.Println(Pokemon.Name + " escaped!")
	return nil

}
