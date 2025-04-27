package api

import (
	"errors"
	"fmt"
)

func Inspect(pokemonName string) error {
	if pokemon, found := catchedPokemons[pokemonName]; !found {
		return errors.New("you have not caught that pokemon")
	} else {
		var hp, attack, defense, specialAtk, specialDef, speed int
		for _, stat := range pokemon.Stats {
			switch stat.Stat.Name {
			case "hp":
				hp = stat.BaseStat
			case "attack":
				attack = stat.BaseStat
			case "defense":
				defense = stat.BaseStat
			case "special-attack":
				specialAtk = stat.BaseStat
			case "special-defense":
				specialDef = stat.BaseStat
			case "spped":
				speed = stat.BaseStat
			}
		}

		var types string = ""
		for _, t := range pokemon.Types {
			types += fmt.Sprintf("\t- %s\n", t.Type.Name)
		}

		fmt.Printf(`Name: %s,
Height: %v
Weight: %v
Stats:
	-hp: %v
	-attack: %v
	-defense: %v
	-special-attack: %v
	-special-defense: %v
	-speed: %v
Types:
%s`, pokemon.Name, pokemon.Height, pokemon.Weight, hp, attack, defense, specialAtk, specialDef, speed, types)
	}

	return nil
}
