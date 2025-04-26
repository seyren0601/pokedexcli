package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type exploreAPIResponse struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel        int           `json:"min_level"`
				MaxLevel        int           `json:"max_level"`
				ConditionValues []interface{} `json:"condition_values"`
				Chance          int           `json:"chance"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetLocationPokemons(location string) ([]string, error) {
	url := BASE_URL + "/location/" + location
	var data []byte
	data, found := cache.Get(url) // Find data in cache first
	if !found {
		res, err := http.Get(url) // fetch (get) data from Url
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body) // Read data from response body as []byte
		if err != nil {
			return nil, err
		}
		cache.Add(url, data) // Cache fetched data
	}

	var apiStruct exploreAPIResponse
	err := json.Unmarshal(data, &apiStruct)
	if err != nil {
		return nil, err
	}

	pokemonList := []string{}
	for _, encounter := range apiStruct.PokemonEncounters {
		pokemonList = append(pokemonList, encounter.Pokemon.Name)
	}

	return pokemonList, nil
}
