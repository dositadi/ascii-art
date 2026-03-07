package config

import (
	m "ascii/pkg/models"
	h "ascii/pkg/utils"
	"strings"
	"unicode"
)

type Artist interface {
	GetLowerCaseStartLine(char rune) int               // Function that handles drawing of lower case alphabets --> done
	GetUpperCaseStartLine(char rune) int               // Function that handles drawing of upper case alphabets --> done
	GetNumberStartLine(char rune) int                  // Function that handles drawing of numbers --> done
	GetSpecialCharStartLine(char rune) int             // Function that handles drawing of Special characters --> done
	ReadFileByLine(startLine int) ([]string, *m.Error) // Function to read specific lines of a file --> done
	CheckIfNewlineAndSplit(input string) []string      // This function checks if the string has a new line character and splits the string with the delimiter --> done
	PrintArtOut(input [][][]string)                    // This function prints out the output to the terminal
}

// Struct for the Sketcher accepting a Setch that implements Artist
type Sketcher struct {
	Sketch Artist
	Input  string
}

// Constructor to create a new sketcher
func NewSketcher(sketcher Artist, input string) *Sketcher {
	return &Sketcher{
		Sketch: sketcher,
		Input:  input,
	}
}

// This is a sketcher function that initiates the sketch of character
func (s *Sketcher) Draw() {
	inputs := s.Sketch.CheckIfNewlineAndSplit(s.Input)

	var maxStore [][][]string

	maxStore = s.DesignAllCharacters(inputs)

	s.Sketch.PrintArtOut(maxStore)
}

func (s *Sketcher) DesignAllCharacters(inputs []string) [][][]string {
	var maxStore [][][]string

	for i := 0; i < len(inputs); i++ {
		currentInput := inputs[i]
		lines := [][]string{}

		if currentInput != "" {
			for _, rn := range currentInput {
				if unicode.IsLetter(rn) {
					if unicode.IsLower(rn) {
						startLine := s.Sketch.GetLowerCaseStartLine(rn)
						result, err := s.Sketch.ReadFileByLine(startLine)
						if err != nil {
							h.PrintError(*err)
							return nil
						}
						lines = append(lines, result)
					} else if unicode.IsUpper(rn) {
						startLine := s.Sketch.GetUpperCaseStartLine(rn)
						result, err := s.Sketch.ReadFileByLine(startLine)
						if err != nil {
							h.PrintError(*err)
							return nil
						}
						lines = append(lines, result)
					}
				} else if unicode.IsDigit(rn) {
					startLine := s.Sketch.GetNumberStartLine(rn)
					result, err := s.Sketch.ReadFileByLine(startLine)
					if err != nil {
						h.PrintError(*err)
						return nil
					}
					lines = append(lines, result)
				} else if strings.ContainsAny(string(rn), " !\"#$%&()*+'-./:;<=>?@[]\\^`{}|~_,") {
					startLine := s.Sketch.GetSpecialCharStartLine(rn)
					result, err := s.Sketch.ReadFileByLine(startLine)
					if err != nil {
						h.PrintError(*err)
						return nil
					}
					lines = append(lines, result)
					continue
				}
			}
			maxStore = append(maxStore, lines)
		} else {
			lines = append(lines, []string{string(rune('\n'))})
			maxStore = append(maxStore, lines)
		}
	}
	return maxStore
}

// This function allows Sketcher to begin its skill
func (s *Sketcher) SkillIt() {
	s.Draw()
}
