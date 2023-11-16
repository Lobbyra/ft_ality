package ftality

import (
	"bufio"
	"ft_ality/src/lib/ftsdl"
	"ft_ality/src/lib/parsing"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

func FtAlity(gramFilePath string) error {
	// OPEN GRAMMAR FILE
	gramFile, errOpenGram := os.Open(gramFilePath)
	if errOpenGram != nil {
		return errOpenGram
	}
	defer gramFile.Close()
	fileScanner := bufio.NewScanner(gramFile)
	// PARSE GRAMMAR FILE
	actionsMap, comboMap, comboSet, errParsing := parsing.Parse(fileScanner)
	if errParsing != nil {
		return errParsing
	}
	// SDL INIT
	window, errInitSDL := ftsdl.InitSDL()
	if errInitSDL != nil {
		return errInitSDL
	}
	// ACTUALLY START TRAINING
	FIGHT(actionsMap, comboMap, comboSet)
	// CLEANING THINGS
	window.Destroy()
	defer sdl.Quit()
	return nil
}
