package arthandlers

import (
	m "ascii/pkg/models"
	h "ascii/pkg/utils"
	"bufio"
	"os"
)

func (s *Sketcher) ReadFileByLine(startLine int) ([]string, *m.Error) {
	file, err := os.Open("/home/gamp/L2E Fellowship/piscine-prompt/ascii-art/asset/model.txt")
	if err != nil {
		return nil, &m.Error{
			Error:  h.SERVER_ERR,
			Detail: h.SERVER_ERR_DETAIL,
		}
	}

	defer file.Close()

	var asciiArt []string

	scanner := bufio.NewScanner(file)

	currentLine := 1

	stopLine := startLine + 8

	for scanner.Scan() {
		if currentLine >= startLine && currentLine <= stopLine {
			asciiArt = append(asciiArt, scanner.Text())
		} else if currentLine > stopLine {
			break
		}
		currentLine++
	}

	if err := scanner.Err(); err != nil {
		return nil, &m.Error{
			Error:  h.SERVER_ERR,
			Detail: h.SERVER_ERR_DETAIL,
		}
	}
	return asciiArt, nil
}
