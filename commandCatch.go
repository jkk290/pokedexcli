package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(pokemonName string, config *Config) error {
	fullUrl := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
	pokemon, err := config.Client.GetPokemon(fullUrl)
	if err != nil {
		return fmt.Errorf("error with catch command: %v", err)
	}
	catchRate := rand.Intn(100) * 2
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if catchRate < pokemon.BaseExperience {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}
	fmt.Printf("%s was caught!\n", pokemonName)
	config.pokeDex[pokemonName] = pokemon
	return nil
}
