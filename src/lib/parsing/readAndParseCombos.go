package parsing

import (
	"bufio"
	"errors"
	"fmt"
	"ft_ality/src/lib"
	"strings"

	"github.com/benbjohnson/immutable"
)

func combosParseErrorOverloading(msg string) error {
	finalMsg := "Combos parsing : " + msg
	return errors.New(finalMsg)
}

func logComboRead(combos []string, comboSentence string) {
	println(
		fmt.Sprintf(
			"%s -> %s",
			lib.StringsToString(combos, "", ","),
			comboSentence,
		),
	)
}

// RETURNS Map[mergedActionNames, combo messages]
func ReadAndParseCombos(
	fileScanner *bufio.Scanner,
	currMap *immutable.Map[string, []string],
	actionNames []string,
	comboSet immutable.Set[string],
) (*immutable.Map[string, []string], immutable.Set[string], error) {
	// READ A FILE LINE
	fileScanner.Scan()
	currLine := fileScanner.Text()
	// NO COMBO TEST
	if len(currLine) == 0 && currMap.Len() == 0 {
		return (nil), (comboSet), (combosParseErrorOverloading(
			"No combo list found"))
	}
	// END CONDITION
	if len(currLine) == 0 {
		return (currMap), (comboSet), (nil)
	}
	// TEST THE LINE FORMAT
	lineSplitted := strings.Split(currLine, ":")
	if len(lineSplitted) != 2 {
		return (nil), (comboSet), (combosParseErrorOverloading(fmt.Sprintf(
			"Line : [%s] : Bad format (expected: [action list]:[combo name])",
			currLine,
		)))
	}
	// TEST IF ALL ACTIONS IN COMBO EXISTS
	combosSplited := strings.Split(lineSplitted[0], ",")
	if checkActionsExistence(combosSplited, actionNames) == false {
		return (nil), (comboSet), (combosParseErrorOverloading(fmt.Sprintf(
			"Unknown action in [%s]",
			currLine,
		)))
	}
	// PRINTING CURRENT COMBO LINE IN HUMAN FORM
	logComboRead(combosSplited, lineSplitted[1])
	combosMerged := lib.StringsToString(combosSplited, "", "")
	currCombos, found := currMap.Get(combosMerged)
	if found == true {
		return ReadAndParseCombos(
			fileScanner,
			currMap.Set(
				combosMerged,
				append(
					currCombos,
					[]string{lineSplitted[1]}...,
				),
			),
			actionNames,
			comboSet.Add(combosMerged),
		)
	} else {
		return ReadAndParseCombos(
			fileScanner,
			currMap.Set(combosMerged, []string{lineSplitted[1]}),
			actionNames,
			comboSet.Add(combosMerged),
		)
	}
}
