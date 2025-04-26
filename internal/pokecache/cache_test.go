package pokecache_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/seyren0601/pokedexcli/internal/pokecache"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second

	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "https://case1.com",
			value: []byte("case1"),
		},
		{
			key:   "http://case2.com",
			value: []byte("case2"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v: ", i), func(t *testing.T) {
			cache := pokecache.NewCache(interval)
			cache.Add(c.key, c.value)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}

			if string(val) != string(c.value) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := pokecache.NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
