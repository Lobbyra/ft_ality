package parsing

import (
	"bufio"
	"fmt"
	"ft_ality/src/lib"
	"strings"

	"github.com/benbjohnson/immutable"
)

// RETURNS Map[mergedActionNames, combo messages]
func ReadAndParseCombos(
	fileScanner *bufio.Scanner,
	currMap *immutable.Map[string, []string],
	actionNames []string,
	comboSet immutable.Set[string],
) (*immutable.Map[string, []string], immutable.Set[string]) {
	fileScanner.Scan()
	currLine := fileScanner.Text()
	if len(currLine) == 0 && currMap.Len() == 0 {
		panic("No combo list found")
	} else if len(currLine) == 0 {
		return (currMap), (comboSet)
	}
	lineSplitted := strings.Split(currLine, ":")
	if len(lineSplitted) != 2 {
		panic(fmt.Sprintf("Line : [%s] : Bad format (expected: [action list]:[combo name])", currLine))
	}
	combosSplited := strings.Split(lineSplitted[0], ",")
	if checkActionsExistence(combosSplited, actionNames) == false {
		panic(fmt.Sprintf("Unknown action in [%s]", currLine))
	}
	println(fmt.Sprintf("%s -> %s", lib.StringsToString(combosSplited, "", ","), lineSplitted[1]))
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
