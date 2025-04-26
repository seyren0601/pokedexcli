package api_test

import (
	"testing"

	"github.com/seyren0601/pokedexcli/api"
)

func TestCatch(t *testing.T) {
	input := []string{"pikachu", "squirtle"}
	output := []string{"pikachu", "squirtle"}
	catched := []string{}

	for _, name := range input {
		for { // Retry until success catch
			if success, _ := api.Catch(name); success {
				catched = append(catched, name)
				break
			}
		}
	}

	if len(output) != len(catched) {
		t.Errorf("Test case failed.\nExpected catches: %v.\nGot %v catches\n", len(output), len(catched))
	}
}

func TestCatchInvalidNames(t *testing.T) {
	input := []string{"abc", "xyz"}
	output := []string{}
	catched := []string{}

	for _, name := range input {
		if success, _ := api.Catch(name); success {
			catched = append(catched, name)
		}
	}

	if len(output) != len(catched) {
		t.Errorf("Test case failed.\nExpected catches: %v.\nGot %v catches\n", len(output), len(catched))
	}
}
