package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type mapArea struct {
	Name string
	Url  string
}

type mapAreaAPIResponse struct {
	Count    int
	Next     string
	Previous string
	Results  []mapArea
}

type config struct {
	Next     string
	Previous string
}

var Cfg config = config{
	Next: "https://pokeapi.co/api/v2/location-area/",
}

func GetMapAreas(url string) ([]mapArea, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var content mapAreaAPIResponse
	err = json.Unmarshal(data, &content)
	if err != nil {
		return nil, err
	}

	Cfg.Previous = content.Previous
	Cfg.Next = content.Next

	return content.Results, nil
}
