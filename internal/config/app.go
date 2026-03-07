package config

import (
	s "ascii/internal/art_handlers"
	m "ascii/pkg/models"
	h "ascii/pkg/utils"
	"os"
)

type App struct {
	Sketcher Sketcher
}

func (a *App) GetInput() (string, *m.Error) {
	if len(os.Args) < 2 {
		return "", &m.Error{
			Error:  h.ERR,
			Detail: h.ERRDETAIL,
		}
	}

	args := os.Args[1:]

	if len(args) < 1 {
		return "", &m.Error{
			Error:  h.MULTI_INPUT_ERR,
			Detail: h.MULTI_INPUT_ERR_DETAIL,
		}
	}

	input := args[0]
	return input, nil
}

func (a *App) Run() {
	sketcher := s.CreateNewSketcher()

	input, err := a.GetInput()
	if err != nil {
		h.PrintError(*err)
	}

	artist := NewSketcher(sketcher, input)

	artist.SkillIt()
}
