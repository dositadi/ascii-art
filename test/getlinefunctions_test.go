package test

import (
	a "ascii/internal/art_handlers"
	"testing"
)

func TestGetLineFunctions(t *testing.T) {
	tests := []struct {
		Input    any
		Expected any
	}{
		{Input: a.CreateNewSketcher().GetLowerCaseStartLine('d'), Expected: 614},
		{Input: a.CreateNewSketcher().GetLowerCaseStartLine('z'), Expected: 812},
		{Input: a.CreateNewSketcher().GetNumberStartLine('9'), Expected: 227},
		{Input: a.CreateNewSketcher().GetNumberStartLine('5'), Expected: 191},
		{Input: a.CreateNewSketcher().GetSpecialCharStartLine('{'), Expected: 821},
		{Input: a.CreateNewSketcher().GetSpecialCharStartLine('|'), Expected: 830},
		{Input: a.CreateNewSketcher().GetUpperCaseStartLine('G'), Expected: 353},
		{Input: a.CreateNewSketcher().GetUpperCaseStartLine('J'), Expected: 380},
	}

	for i, cases := range tests {
		if cases.Input != cases.Expected {
			t.Errorf("Case %v's Output: %v, does not equal its Expected: %v", i, cases.Input, cases.Expected)
		}
	}
}
