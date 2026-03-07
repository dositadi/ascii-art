package test

import (
	a "ascii/internal/art_handlers"
	"slices"
	"testing"
)

func TestCheckNewLine(t *testing.T) {
	tests := []struct {
		Input    []string
		Expected []string
	}{
		{Input: a.CreateNewSketcher().CheckIfNewlineAndSplit("hello\nworld"), Expected: []string{"hello", "world"}},
		{Input: a.CreateNewSketcher().CheckIfNewlineAndSplit("hello\nworld\ngoodday"), Expected: []string{"hello", "world", "goodday"}},
		{Input: a.CreateNewSketcher().CheckIfNewlineAndSplit("excess\n\nnewline"), Expected: []string{"excess", "", "newline"}},
	}

	for i, cases := range tests {
		if slices.Compare(cases.Input, cases.Expected) != 0 {
			t.Errorf("Case %v's Output: %+v, does not equal its Expected: %+v", i, cases.Input, cases.Expected)
		}
	}
}
