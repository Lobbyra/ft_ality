package ftality

import (
	"fmt"
	"ft_ality/src/lib"
	"ft_ality/src/lib/parsing"
	"strings"

	"github.com/benbjohnson/immutable"
	"github.com/fatih/color"
	"github.com/veandco/go-sdl2/sdl"
)

func printSuccessCombo(currentCombo string, combosFound []string, layer int) {
	fmt.Printf("\r%s\r", strings.Repeat(" ", len(currentCombo)))
	color.Red("🔥 COMBO x %d !!!", layer+1)
	println(lib.StringsToString(combosFound, "", "\n"))
	print("\n")
	fmt.Printf("%s", currentCombo)
}

/*
 * This is the recursive function that will stack combos and print them.
 *
 * If a combo ends or is not found, it will end and the program
 * goes back to run().
 */
func Try(
	actionsMap *immutable.Map[string, string],
	comboMap *immutable.Map[string, []string],
	comboSet immutable.Set[string],
	currentCombo string,
	layer int,
) {
	fmt.Printf("\r%s", currentCombo)
	// DISPLAY THE CURRENT COMBO IF EXIST
	combosFound, comboFound := comboMap.Get(currentCombo)
	if comboFound == true {
		printSuccessCombo(currentCombo, combosFound, layer)
	}
	if parsing.TestComboEnded(currentCombo, *comboSet.Iterator()) == false {
		if layer == 0 {
			print(" 👉 (no combo found)\n\n")
		} else {
			print(" ⬇️  (combo finished)\n\n")
		}
		return
	}
	// WAIT FOR A VALID KEYBOARD
	for {
		event := sdl.PollEvent()
		if event != nil && event.GetType() == sdl.KEYDOWN {
			// GET PRESSED KEY NAME
			key := sdl.GetKeyName((event.(*sdl.KeyboardEvent).Keysym.Sym))
			// FIND THE ACTION OF THE KEYPRESSED
			action, actionFound := actionsMap.Get(key)
			if actionFound == true {
				Try(
					actionsMap,
					comboMap,
					comboSet,
					currentCombo+action,
					layer+1,
				)
			} else {
				color.Magenta(
					"WARNING: action : %s -> not found in mapping, check the mapping",
					key,
				)
			}
			break
		}
	}
}
