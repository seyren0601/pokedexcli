package helpers

import "strings"

func CleanInput(text string) []string {
	var tokens []string = strings.Split(text, " ")
	for i := 0; i < len(tokens); i++ {
		tokens[i] = strings.ToLower(strings.Trim(tokens[i], " "))
	}

	return tokens
}
