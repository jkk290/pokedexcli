package main

import "fmt"

func commandHelp(str string, config *Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Println(cmd.name + ": " + cmd.description)
	}
	fmt.Println()
	return nil
}
