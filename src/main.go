package main

import (
	"bufio"
	"fmt"
	"ft_ality/src/lib"
	"ft_ality/src/lib/parsing"
	"os"

	"github.com/benbjohnson/immutable"
	"github.com/fatih/color"
	"github.com/veandco/go-sdl2/sdl"
)

func printUsage() {
	println("USAGE:\n    ./ft_ality synthax-file")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func try(
	actionsMap *immutable.Map[string, string],
	comboMap *immutable.Map[string, []string],
	currentCombo string,
) {
	fmt.Printf("\r%s", currentCombo)
	// DISPLAY THE CURRENT COMBO IF EXIST
	combo, comboFound := comboMap.Get(currentCombo)
	if comboFound == true {
		print("\n\n")
		color.Red("COMBO !!!")
		fmt.Printf("%s\n", currentCombo)
		println(lib.StringsToString(combo, "", "\n"))
		print("\n")
	}
	for {
		event := sdl.PollEvent()
		if event != nil && event.GetType() == sdl.KEYDOWN {
			keyEvent := event.(*sdl.KeyboardEvent)
			key := sdl.GetKeyName(keyEvent.Keysym.Sym)
			// LOOKING FOR
			action, actionFound := actionsMap.Get(key)
			if actionFound == true {
				try(actionsMap, comboMap, currentCombo+action)
			}
			break
		}
	}
}

func run(
	actionsMap *immutable.Map[string, string],
	comboMap *immutable.Map[string, []string],
	currentCombo string,
) {
	for {
		event := sdl.PollEvent()
		if event != nil && event.GetType() == sdl.KEYDOWN {
			keyEvent := event.(*sdl.KeyboardEvent)
			key := sdl.GetKeyName(keyEvent.Keysym.Sym)
			action, actionFound := actionsMap.Get(key)
			if actionFound == true {
				try(actionsMap, comboMap, currentCombo+action)
				println()
			}
		}
	}
}

func ftAlity(gramFilePath string) {
	gramFile, err := os.Open(gramFilePath)
	switch err {
	case nil:
		defer gramFile.Close()
		fileScanner := bufio.NewScanner(gramFile)
		// PARSE KEY MAPPINGS
		println("Key mappings:")
		actionsMap, actionNames := parsing.ReadAndParseAction(
			fileScanner,
			immutable.NewMap[string, string](nil), []string{},
		)
		// PARSE THE COMBO LIST
		println("\nCombos:")
		comboMap := parsing.ReadAndParseCombos(
			fileScanner,
			immutable.NewMap[string, []string](nil), actionNames,
		)
		// SDL INIT
		if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
			panic(err)
		}
		defer sdl.Quit()
		println("----------------------\n")
		run(actionsMap, comboMap, "")
		println("End of the training...")
	default:
		panic(err)
	}
}

func main() {
	switch len(os.Args) {
	case 2:
		ftAlity(os.Args[1])
		break
	default:
		printUsage()
	}
}
