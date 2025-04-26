package api

import (
	"time"

	"github.com/seyren0601/pokedexcli/internal/pokecache"
)

var Cfg config = config{
	Next: BASE_URL + "/location",
}
var cache pokecache.Cache = pokecache.NewCache(5 * time.Second)
var catchedPokemons map[string]Pokemon = map[string]Pokemon{}
