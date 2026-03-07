package arthandlers

import (
	h "ascii/pkg/utils"
)

func (s *Sketcher) GetLowerCaseStartLine(char rune) int {
	if lineNumber, ok := h.LowerCaseAlphabetsArt[char]; ok {
		return lineNumber
	}
	return -1
}
