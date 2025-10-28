package main

import "fmt"

func commandPokedex(str string, config *Config) error {
	if len(config.pokeDex) == 0 {
		fmt.Println("Your Pokedex is empty, go catch some Pokemon!")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for name, _ := range config.pokeDex {
		fmt.Printf("- %s\n", name)
	}
	return nil
}
