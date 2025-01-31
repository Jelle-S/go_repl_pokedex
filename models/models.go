package models

import (
	"github.com/Jelle-S/pokedexcli/internal/pokecache"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(config *ConfigType, arguments []string) error
}

type ConfigType struct {
	Next     string
	Previous string
	Cache    *pokecache.Cache
	Pokedex  map[string]Pokemon
}

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name              string             `json:"name"`
	URL               string             `json:"url"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name    string `json:"name"`
	BaseExp int    `json:"base_experience"`
}
