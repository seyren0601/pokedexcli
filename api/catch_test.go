package api_test

import (
	"fmt"
	"testing"

	"github.com/seyren0601/pokedexcli/api"
)

func TestCatch(t *testing.T) {
	cases := []struct {
		input  []string
		output []string
	}{
		{
			[]string{"pikachu", "squirtle"},
			[]string{"pikachu", "squirtle"},
		},
		{
			[]string{"abc", "xyz"},
			[]string{},
		},
	}

	for _, c := range cases {
		testname := fmt.Sprintf("%v", c.input)
		t.Run(testname, func(t *testing.T) {
			var catched []string = []string{}
			for _, name := range c.input {
				for { // Retry until success catch or
					success, err := api.Catch(name)
					if err != nil { // Exception
						break
					}

					if success { // Success catch
						catched = append(catched, name)
						break
					}
				}
			}
			if len(c.output) != len(catched) {
				t.Errorf("Test case failed.\nExpected catches: %v.\nGot %v catches\n", len(c.output), len(catched))
			}
		})
	}

}
