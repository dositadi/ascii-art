package arthandlers

import (
	h "ascii/pkg/utils"
)

func (s *Sketcher) GetUpperCaseStartLine(char rune) int {
	if lineNumber, ok := h.UpperCaseAlphabetsArt[char]; ok {
		return lineNumber
	}
	return -1
}
