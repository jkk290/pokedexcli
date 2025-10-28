package main

import (
	"fmt"

	"github.com/jkk290/pokedexcli/internal/pokeapi"
)

func commandMapb(config *Config) error {
	if config.Previous == "" {
		fmt.Println("No previous locations to display")
		return nil
	}

	url := config.Previous

	locations, err := pokeapi.GetLocations(url)
	if err != nil {
		return fmt.Errorf("error with map command: %v", err)
	}

	for _, location := range locations.Results {
		fmt.Println(*location.Name)
	}

	config.Next = *locations.Next

	if locations.Previous != nil {
		config.Previous = *locations.Previous
	} else {
		config.Previous = ""
	}

	return nil
}
