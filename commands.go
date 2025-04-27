package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/seyren0601/pokedexcli/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func([]string) error
}

var commands map[string]cliCommand

func InitCommandList() {
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
		"explore": {
			name:        "explore",
			description: "Displays the pokemon list of given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Tries to catch a specified pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a registered pokemon",
			callback:    commandInspect,
		},
	}
}

func commandExit(parameters []string) error {
	if len(parameters) > 0 {
		return errors.New("this command doesn't take parameters")
	}

	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)

	return nil
}

func commandHelp(parameters []string) error {
	if len(parameters) > 0 {
		return errors.New("this command doesn't take parameters")
	}

	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf(`Usage:
	
`)
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandMap(parameters []string) error {
	if len(parameters) > 0 {
		return errors.New("this command doesn't take parameters")
	}

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

func commandMapb(parameters []string) error {
	if len(parameters) > 0 {
		return errors.New("this command doesn't take parameters")
	}

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

func commandExplore(parameters []string) error {
	if len(parameters) != 1 {
		return errors.New("this command takes 1 parameter")
	}
	location := parameters[0]

	pokemons, err := api.GetLocationPokemons(location)
	if err != nil {
		return err
	}

	for _, pokemon := range pokemons {
		fmt.Printf("%s\n", pokemon)
	}

	return nil
}

func commandCatch(parameters []string) error {
	if len(parameters) != 1 {
		return errors.New("this command takes 1 parameter")
	}
	pokemonName := parameters[0]

	successful, err := api.Catch(pokemonName)
	if err != nil {
		return err
	}

	if successful {
		fmt.Printf("%s was caught!\n", pokemonName)
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

func commandInspect(parameters []string) error {
	if len(parameters) != 1 {
		return errors.New("this command takes 1 parameter")
	}
	pokemonName := parameters[0]

	err := api.Inspect(pokemonName)
	if err != nil {
		return err
	}

	return nil
}
