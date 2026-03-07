package arthandlers

import (
	h "ascii/pkg/utils"
)

func (s *Sketcher) GetSpecialCharStartLine(char rune) int {
	if lineNumber, ok := h.SpecialCharactersArt[char]; ok {
		return lineNumber
	}
	return -1
}
