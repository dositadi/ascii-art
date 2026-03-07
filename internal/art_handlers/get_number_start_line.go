package arthandlers

import (
	h "ascii/pkg/utils"
)

func (s *Sketcher) GetNumberStartLine(char rune) int {
	if lineNumber, ok := h.NumbersArt[char]; ok {
		return lineNumber
	}
	return -1
}
