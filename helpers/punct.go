package helpers

import (
	"regexp"
)

func Punct(text string) string {

	punc := regexp.MustCompile(`\s+([.,!?;:])`)
	text = punc.ReplaceAllString(text, "$1")

	punc_space := regexp.MustCompile(`([.,!?;:])([a-zA-Z0-9])`)
	text = punc_space.ReplaceAllString(text, "$1 $2")

	return text
}
