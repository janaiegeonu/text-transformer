package helpers

import "strings"

func Compiler(text string) string {

	lines := strings.Split(text, "\n")

	for i := range lines {
		lines[i] = Quotes(lines[i])
		lines[i] = Casing(lines[i])
		lines[i] = Articles(lines[i])
		lines[i] = Base(lines[i])
		lines[i] = Punct(lines[i])
	}

	return strings.Join(lines, "\n")
}
