package main

import "fmt"

func commandInspect(pokemonName string, config *Config) error {
	pokemon, exists := config.pokeDex[pokemonName]
	if exists {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("-%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pType := range pokemon.Types {
			fmt.Printf("- %s\n", pType.Type.Name)
		}
		return nil
	}
	return fmt.Errorf("%s not found in pokedex", pokemonName)

}
