package main

import (
	"bufio"
	"fmt"
	"ft_ality/src/lib"
	"ft_ality/src/lib/parsing"
	"os"
	"strings"

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
	comboSet immutable.Set[string],
	currentCombo string,
	layer int,
) {
	fmt.Printf("\r%s", currentCombo)
	// DISPLAY THE CURRENT COMBO IF EXIST
	combo, comboFound := comboMap.Get(currentCombo)
	if comboFound == true {
		fmt.Printf("\r%s\r", strings.Repeat(" ", len(currentCombo)))
		color.Red("🔥 COMBO !!!")
		println(lib.StringsToString(combo, "", "\n"))
		print("\n")
		fmt.Printf("%s", currentCombo)
	}
	if parsing.TestComboEnded(currentCombo, *comboSet.Iterator()) == false {
		if layer == 0 {
			print(" 👉 (no combo found)")
		} else {
			print(" ⬇️  (combo finished)")
		}
		print("\n\n")
		return
	}
	for {
		event := sdl.PollEvent()
		if event != nil && event.GetType() == sdl.KEYDOWN {
			key := sdl.GetKeyName((event.(*sdl.KeyboardEvent).Keysym.Sym))
			action, actionFound := actionsMap.Get(key)
			if actionFound == true {
				try(actionsMap, comboMap, comboSet, currentCombo+action, layer+1)
			}
			break
		}
	}
}

func run(
	actionsMap *immutable.Map[string, string],
	comboMap *immutable.Map[string, []string],
	comboSet immutable.Set[string],
) {
	for {
		event := sdl.PollEvent()
		if event != nil && event.GetType() == sdl.KEYDOWN {
			key := sdl.GetKeyName((event.(*sdl.KeyboardEvent).Keysym.Sym))
			action, actionFound := actionsMap.Get(key)
			if actionFound == true {
				try(actionsMap, comboMap, comboSet, action, 0)
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
		comboMap, comboSet := parsing.ReadAndParseCombos(
			fileScanner,
			immutable.NewMap[string, []string](nil), actionNames,
			immutable.NewSet[string](nil),
		)
		// SDL INIT
		if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
			panic(err)
		}
		defer sdl.Quit()
		println("----------------------\n")
		run(actionsMap, comboMap, comboSet)
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
