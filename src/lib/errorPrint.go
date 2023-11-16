package lib

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func ErrorPrint(err error) {
	fmt.Fprintf(os.Stderr, "💥 ")
	red := color.New(color.FgRed)
	red.Fprint(os.Stderr, err.Error())
}
