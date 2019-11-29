package main

import (
	"strings"

	"github.com/gookit/color"
)

func commandNotFound(command string) {
	stringLen := len(command)
	highlightThing := strings.Repeat("^", stringLen)
	color.FgBlue.Print("1 | ")
	color.FgRed.Println("gosh: " + command + "")
	color.FgBlue.Print("  |   ")
	color.FgRed.Println("    " + highlightThing + " command not found")
}
func directoryNotFound(dir string) {
	stringLen := len(dir)
	highlightThing := strings.Repeat("^", stringLen)
	color.FgBlue.Print("1 | ")
	color.FgRed.Println("gosh: " + dir + "")
	color.FgBlue.Print("  |   ")
	color.FgRed.Println("    " + highlightThing + " directory not found")
}
