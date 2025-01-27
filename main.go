package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Jelle-S/pokedexcli/internal/pokecache"

	"github.com/Jelle-S/pokedexcli/models"
)

func main() {
	commands := supportedCommands()
	scanner := bufio.NewScanner(os.Stdin)
	config := models.ConfigType{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
		Cache:    pokecache.NewCache(5 * time.Second),
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		inputs := cleanInput(scanner.Text())
		c := inputs[0]
		command, ok := commands[c]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.Callback(&config, inputs[1:])
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
