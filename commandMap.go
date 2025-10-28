package main

import (
	"fmt"
)

func commandMap(config *Config) error {
	var url string

	if config.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	} else {
		url = config.Next
	}

	locations, err := config.Client.GetLocations(url)
	if err != nil {
		return fmt.Errorf("error with map command: %v", err)
	}

	for _, location := range locations.Results {
		fmt.Println(*location.Name)
	}

	config.Next = *locations.Next
	if locations.Previous != nil {
		config.Previous = *locations.Previous
	}

	return nil

}
