package helpers

import (
	"strings"
)

func Articles(text string) string {

	words := strings.Fields(text)
	vowel := "aeiouhAEIOUH"

	for i := 0; i < len(words); i++ {

		if i+1 < len(words) {

			nextwords := words[i+1][:1]

			switch words[i] {
			case "A":
				if strings.ContainsAny(nextwords, vowel) {
					words[i] = "An"
				}
			case "a":
				if strings.ContainsAny(nextwords, vowel) {
					words[i] = "an"
				}
			case "An":
				if !strings.ContainsAny(nextwords, vowel) {
					words[i] = "A"
				}
			case "an":
				if !strings.ContainsAny(nextwords, vowel) {
					words[i] = "a"
				}
			}

		}

	}
	return strings.Join(words, " ")

}
