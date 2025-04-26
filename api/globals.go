package api

import (
	"time"

	"github.com/seyren0601/pokedexcli/internal/pokecache"
)

var Cfg config = config{
	Next: BASE_URL,
}
var cache pokecache.Cache = pokecache.NewCache(5 * time.Second)
