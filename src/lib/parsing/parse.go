package parsing

import (
	"bufio"
	"errors"

	"github.com/benbjohnson/immutable"
)

func Parse(
	fileScanner *bufio.Scanner,
) (
	*immutable.Map[string, string],
	*immutable.Map[string, []string],
	immutable.Set[string],
	error,
) {
	// PARSE KEY MAPPINGS
	println("Key mappings:")
	actionsMap, actionNames, errParsAction := ReadAndParseAction(
		fileScanner,
		immutable.NewMap[string, string](nil),
		[]string{},
	)
	if errParsAction != nil {
		return (nil), (nil), (immutable.NewSet[string](nil)), (errParsAction)
	} else if actionsMap.Len() == 0 {
		return (nil),
			(nil),
			(immutable.NewSet[string](nil)), (errors.New("No keybinding found"))
	}
	// PARSE THE COMBO LIST
	println("\nCombos:")
	comboMap, comboSet, errParsCombo := ReadAndParseCombos(
		fileScanner,
		immutable.NewMap[string, []string](nil),
		actionNames,
		immutable.NewSet[string](nil),
	)
	if errParsCombo != nil {
		return (nil), (nil), (immutable.NewSet[string](nil)), (errParsCombo)
	}
	return (actionsMap), (comboMap), (comboSet), (nil)
}
