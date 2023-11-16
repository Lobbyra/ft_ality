package ftality

import (
	"github.com/benbjohnson/immutable"
	"github.com/fatih/color"
	"github.com/veandco/go-sdl2/sdl"
)

// This function is the root of the training -> the "infinite" loop
func FIGHT(
	actionsMap *immutable.Map[string, string],
	comboMap *immutable.Map[string, []string],
	comboSet immutable.Set[string],
) {
	println("----------------------")
	color.Green("🏃‍♀️ Start of the training...\n\n")
	// WAIT FOR A VALID KEYBOARD
	for {
		event := sdl.PollEvent()
		if event != nil && event.GetType() == sdl.KEYDOWN {
			key := sdl.GetKeyName((event.(*sdl.KeyboardEvent).Keysym.Sym))
			if key == "Escape" {
				break
			}
			action, actionFound := actionsMap.Get(key)
			if actionFound == true {
				Try(actionsMap, comboMap, comboSet, action, 0)
			} else {
				color.Magenta(
					"WARNING: action : %s -> not found in mapping, check the mapping",
					key,
				)
			}
		}
	}
	color.Green("😮‍💨 End of the training...")
}
