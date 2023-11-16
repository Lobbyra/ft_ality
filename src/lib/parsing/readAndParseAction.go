package parsing

import (
	"bufio"
	"errors"
	"fmt"
	"ft_ality/src/lib"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/benbjohnson/immutable"
)

func actionsParseErrorOverloading(err error) error {
	finalMsg := "Action parsing : " + err.Error()
	return errors.New(finalMsg)
}

func isLineValidFormat(lineSplitted []string) error {
	firstKeyRune, _ := utf8.DecodeRune([]byte(lineSplitted[0]))
	// PARSING GLOBAL LINE
	if len(lineSplitted) != 2 {
		return errors.New("Only a pair of value is attended")
	}
	// PARSING KEY
	if len(lineSplitted[0]) == 0 {
		return errors.New("Missing action key")
	}
	if len(lineSplitted[0]) != 1 &&
		lib.Contains(
			lineSplitted[0],
			[]string{"Left", "Right", "Up", "Down"},
		) == false {
		return errors.New(
			"Non single character key must be an arrow (Left|Right|Up|Down)",
		)
	} else if unicode.IsUpper(firstKeyRune) == false {
		return errors.New("Single key must be uppercase")
	}
	// PARSING NAME
	if len(lineSplitted[1]) == 0 {
		return errors.New("Missing action name")
	}
	return nil
}

// RETURNS Map[keyboard key:action name], actionNames
func ReadAndParseAction(
	fileScanner *bufio.Scanner,
	currMap *immutable.Map[string, string],
	actionNames []string,
) (*immutable.Map[string, string], []string, error) {
	// READ A NEW LINE
	fileScanner.Scan()
	currLine := fileScanner.Text()
	// END CONDITION
	if len(currLine) == 0 {
		return (currMap), (actionNames), (nil)
	}
	lineSplitted := strings.Split(currLine, ":")
	// PARSE THE CURRENT LINE
	errLine := isLineValidFormat(lineSplitted)
	if errLine != nil {
		addLineErr := errors.New(
			fmt.Sprintf("[%s] : ", currLine) + errLine.Error(),
		)
		return (nil), (nil), (actionsParseErrorOverloading(addLineErr))
	}
	// LOG THE PARSED LINE
	fmt.Println(fmt.Sprintf("%s -> %s", lineSplitted[0], lineSplitted[1]))
	// GO TO THE NEXT LINE
	return ReadAndParseAction(
		fileScanner,
		currMap.Set(lineSplitted[0], lineSplitted[1]),
		append(
			actionNames,
			[]string{lineSplitted[1]}...,
		),
	)
}
