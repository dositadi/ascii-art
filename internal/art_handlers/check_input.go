package arthandlers

import (
	"strings"
)

func (s *Sketcher) CheckIfNewlineAndSplit(input string) []string {
	var output []string
	containsNewLine := false

	if strings.Contains(input, "\\n") {
		containsNewLine = true
	}

	if containsNewLine {
		output = strings.Split(input, "\\n")
	} else {
		output = append(output, input)
	}
	return output
}
