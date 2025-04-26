package main

import (
	"testing"

	"github.com/seyren0601/pokedexcli/helpers"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := helpers.CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("mismatched slice length between expected and actual.\nexpected: %v\nactual: %v\n", len(c.expected), len(actual))
		}

		for i := range actual {
			token := actual[i]
			expected := c.expected[i]

			if token != expected {
				t.Errorf("mismatched.\nexpected: %s\nactual: %s\n", expected, token)
			}
		}
	}
}
