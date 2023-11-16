package main

import (
	ftality "ft_ality/src/ftAlity"
	"ft_ality/src/lib"
	"os"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) != 2 {
		color.Red("Wrong number of arguments")
		lib.PrintUsage()
		os.Exit(1)
	} else if os.Args[1] == "-h" || os.Args[1] == "--help" {
		lib.PrintUsage()
	} else {
		errFtAlity := ftality.FtAlity(os.Args[1])
		if errFtAlity != nil {
			lib.ErrorPrint(errFtAlity)
			os.Exit(1)
		}
	}
}
