package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/seyren0601/pokedexcli/api"
	"github.com/seyren0601/pokedexcli/helpers"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)

	return nil
}

func commandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf(`Usage:
	
`)
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap() error {
	if api.Cfg.Next == "" {
		fmt.Println("you're on the last page")
		return nil
	}

	mapAreas, err := api.GetMapAreas(api.Cfg.Next)
	if err != nil {
		return err
	}

	for _, mapArea := range mapAreas {
		fmt.Printf("%s\n", mapArea.Name)
	}

	return nil
}

func commandMapb() error {
	if api.Cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	mapAreas, err := api.GetMapAreas(api.Cfg.Previous)
	if err != nil {
		return err
	}

	for _, mapArea := range mapAreas {
		fmt.Printf("%s\n", mapArea.Name)
	}

	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays current map area page",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Return to previous map page",
			callback:    commandMapb,
		},
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		cleanedInput := helpers.CleanInput(input)
		firstWord := cleanedInput[0]

		if command, ok := commands[firstWord]; !ok {
			fmt.Printf("Unknown command\n")
		} else {
			command.callback()
		}
	}
}
