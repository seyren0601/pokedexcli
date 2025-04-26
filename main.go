package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/seyren0601/pokedexcli/helpers"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	InitCommandList()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		cleanedInput := helpers.CleanInput(input)
		firstWord := cleanedInput[0]

		if command, ok := commands[firstWord]; !ok {
			fmt.Printf("Unknown command\n")
		} else {
			err := command.callback(cleanedInput[1:])
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
