package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jkk290/pokedexcli/internal/pokeapi"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := Config{
		Next:     "",
		Previous: "",
		Client:   pokeapi.NewClient(30),
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cmdInput := scanner.Text()
		cleanedInput := cleanInput(cmdInput)
		if len(cleanedInput) == 0 {
			continue
		}

		userCmd := cleanedInput[0]

		runCmd, exists := getCommands()[userCmd]
		if exists {
			err := runCmd.callback(&config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Command not found")
			continue
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
}

type Config struct {
	Next     string
	Previous string
	Client   *pokeapi.Client
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next list of 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous list of 20 locations",
			callback:    commandMapb,
		},
	}
}
