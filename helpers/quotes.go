package helpers

import (
	"strings"
)

func Quotes(text string) string {
	words := strings.Split(text, "\"")

	for i := range words {
		if i%2 == 1 {
			words[i] = strings.TrimSpace(words[i])
		}
	}

	result := strings.Join(words, "\"")

	single := strings.Split(result, "'")

	for i := range single {
		if i%2 == 1 {
			single[i] = strings.TrimSpace(single[i])
		}
	}
	return strings.Join(single, "'")

}
