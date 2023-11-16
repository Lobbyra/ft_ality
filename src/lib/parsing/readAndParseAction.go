package parsing

import (
	"bufio"
	"fmt"
	"ft_ality/src/lib"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/benbjohnson/immutable"
)

// RETURNS Map[keyboard key:action name], actionNames
func ReadAndParseAction(
	fileScanner *bufio.Scanner,
	currMap *immutable.Map[string, string],
	actionNames []string,
) (*immutable.Map[string, string], []string) {
	fileScanner.Scan()
	currLine := fileScanner.Text()
	if len(currLine) == 0 {
		return (currMap), (actionNames)
	}
	lineSplitted := strings.Split(currLine, ":")
	firstKeyRune, _ := utf8.DecodeRune([]byte(lineSplitted[0]))
	if len(lineSplitted) != 2 ||
		(len(lineSplitted[0]) != 1 && lib.Contains(lineSplitted[0], []string{"Left", "Right", "Up", "Down"}) == false) ||
		unicode.IsUpper(firstKeyRune) == false {
		panic(fmt.Sprintf("Line : [%s] : Bad format (expected: [uppercase letter|'Left'|'Right'|'Up'|'Down'])", currLine))
	}
	fmt.Println(fmt.Sprintf("%s -> %s", lineSplitted[0], lineSplitted[1]))
	return ReadAndParseAction(
		fileScanner,
		currMap.Set(lineSplitted[0], lineSplitted[1]),
		append(
			actionNames,
			[]string{lineSplitted[1]}...,
		),
	)
}
