package main

import (
	"fmt"
)

func commandMapb(str string, config *Config) error {
	if config.Previous == "" {
		fmt.Println("No previous locations to display")
		return nil
	}

	url := config.Previous

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
	} else {
		config.Previous = ""
	}

	return nil
}
