package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cmdInput := scanner.Text()
		cleanedInput := cleanInput(cmdInput)
		if len(cleanedInput) == 0 {
			continue
		}

		fmt.Println("Your command was:", cleanedInput[0])
	}
}
