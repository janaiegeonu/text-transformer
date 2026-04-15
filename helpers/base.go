package helpers

import (
	"strconv"
	"strings"
)

func Base(text string) string {

	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {

		if words[i] == "(bin)" && i > 0 {
			base, err := strconv.ParseInt(words[i-1], 2, 64)
			if err != nil {
				words = append(words[:i], words[i+1:]...)
				i--
			}
			if err == nil {
				words[i-1] = strconv.FormatInt(base, 10)

				words = append(words[:i], words[i+1:]...)
				i--
			}
		}

		if words[i] == "(hex)" && i > 0 {
			base, err := strconv.ParseInt(words[i-1], 16, 64)
			if err != nil {
				words = append(words[:i], words[i+1:]...)
				i--
			}
			if err == nil {
				words[i-1] = strconv.FormatInt(base, 10)

				words = append(words[:i], words[i+1:]...)
				i--
			}
		}
	}
	return strings.Join(words, " ")

}
