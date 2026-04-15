package helpers

import (
	"strconv"
	"strings"
)

func Casing(text string) string {

	words := strings.Fields(text)

	for i := len(words) - 1; i >= 0; i-- {
		// upper casing with index (up, int)
		if strings.HasPrefix(words[i], "(up,") && i+1 < len(words) {
			index := strings.TrimSuffix(words[i+1], ")")
			num, err := strconv.Atoi(index)
			if err != nil {
				words[i-1] = strings.ToUpper(words[i-1])

				words = append(words[:i], words[i+2:]...)
				i--
			}
			if err == nil {
				for j := 1; j <= num; j++ {
					if i-j >= 0 {
						words[i-j] = strings.ToUpper(words[i-j])
					}
				}
				words = append(words[:i], words[i+2:]...)
				i--
			}
		}
		// lower casing with index (low, int)

		if strings.HasPrefix(words[i], "(low,") && i+1 < len(words) {
			index := strings.TrimSuffix(words[i+1], ")")
			num, err := strconv.Atoi(index)
			if err != nil {
				words[i-1] = strings.ToLower(words[i-1])

				words = append(words[:i], words[i+2:]...)
				i--
			}
			if err == nil {
				for j := 1; j <= num; j++ {
					if i-j >= 0 {
						words[i-j] = strings.ToLower(words[i-j])
					}
				}
				words = append(words[:i], words[i+2:]...)
				i--
			}
		}
		// capitalizing casing with index (cap, int)

		if strings.HasPrefix(words[i], "(cap,") && i+1 < len(words) {
			index := strings.TrimSuffix(words[i+1], ")")
			num, err := strconv.Atoi(index)
			if err != nil {
				words[i-1] = strings.ToUpper(words[i-1][:1]) + strings.ToLower(words[i-1][1:])
				words = append(words[:i], words[i+2:]...)
				i--
			}
			if err == nil {
				for j := 1; j <= num; j++ {
					if i-j >= 0 {

						words[i-j] = strings.ToUpper(words[i-j][:1]) + strings.ToLower(words[i-j][1:])
					}
				}
				words = append(words[:i], words[i+2:]...)
				i--
			}
		}

	}

	for i := len(words) - 1; i >= 0; i-- {
		// upper casing (up)
		if words[i] == "(up)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1])

			words = append(words[:i], words[i+1:]...)
			i--
		}

		// lower casing (low)
		if words[i] == "(low)" && i > 0 {
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}

		// capitalization casing (cap)
		if words[i] == "(cap)" && i > 0 {
			words[i-1] = strings.ToUpper(words[i-1][:1]) + strings.ToLower(words[i-1][1:])
			words = append(words[:i], words[i+1:]...)
			i--
		}

		if i == 0 && (words[i] == "(cap)" || words[i] == "(up)" || words[i] == "(low)") {

			words = append(words[:i], words[i+1:]...)
			i--
		}

	}
	return strings.Join(words, " ")
}
