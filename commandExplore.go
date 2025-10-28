package main

import "fmt"

func commandExplore(areaName string, config *Config) error {
	fullUrl := "https://pokeapi.co/api/v2/location-area/" + areaName

	encounters, err := config.Client.GetEncounters(fullUrl)
	if err != nil {
		return fmt.Errorf("error with explore command: %v", err)
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range encounters.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
