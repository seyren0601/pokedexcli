package api

import "fmt"

func Pokedex() error {
	fmt.Println("Your pokedex:")

	var pokemons string
	for _, pokemon := range catchedPokemons {
		pokemons += fmt.Sprintf("\t- %s\n", pokemon.Name)
	}

	fmt.Print(pokemons)

	return nil
}
