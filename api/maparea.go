package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/seyren0601/pokedexcli/internal/pokecache"
)

type MapArea struct {
	Name string
	Url  string
}

type mapAreaAPIResponse struct {
	Count    int
	Next     string
	Previous string
	Results  []MapArea
}

type config struct {
	Next     string
	Previous string
}

var Cfg config = config{
	Next: "https://pokeapi.co/api/v2/location-area/", //
}
var cache pokecache.Cache = pokecache.NewCache(5 * time.Second)

func GetMapAreas(url string) ([]MapArea, error) {
	var data []byte
	data, found := cache.Get(url) // Find data in cache first
	if !found {                   // Can't find data in cache
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

	var content mapAreaAPIResponse
	err := json.Unmarshal(data, &content) // Parse data into struct
	if err != nil {
		return nil, err
	}

	Cfg.Previous = content.Previous // Set previous page url
	Cfg.Next = content.Next         // Set next page url

	return content.Results, nil
}
