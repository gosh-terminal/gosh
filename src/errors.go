package main

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
)

func invalidNumberOfArgs(command string) {
	stringLen := len(command)
	highlightThing := strings.Repeat("^", stringLen)
	color.FgBlue.Print("1 | ")
	color.FgRed.Println("❌  gosh: " + command + "")
	color.FgBlue.Print("  |   ")
	color.FgLightBlue.Print("       " + highlightThing)
	color.FgRed.Print("      Error: ")
	fmt.Println("Invalid number of arguments!")
}
func commandNotFound(command string) {
	stringLen := len(command)
	highlightThing := strings.Repeat("^", stringLen)
	color.FgBlue.Print("1 | ")
	color.FgRed.Println("❌  gosh: " + command + "")
	color.FgBlue.Print("  |   ")
	color.FgLightBlue.Print("       " + highlightThing)
	color.FgRed.Print("      Error: ")
	fmt.Println("Command not found")
}
func directoryNotFound(dir string) {
	stringLen := len(dir)
	highlightThing := strings.Repeat("^", stringLen)
	color.FgBlue.Print("1 | ")
	color.FgRed.Println("❌  gosh: cd " + dir + "")
	color.FgBlue.Print("  |   ")
	color.FgBlue.Print("          " + highlightThing)
	color.FgRed.Print("      Error:")
	fmt.Println("Directory not found")
}
func pipeError(command string) {
	indexOf := strings.Index(command, " > ")
	highlightThing := strings.Repeat("^", len(command[indexOf+3:]))
	spaces := strings.Repeat(" ", len(command[:indexOf])+1)
	color.FgBlue.Print("1 | ")
	color.FgRed.Println("❌  gosh: " + " " + command + "")
	color.FgBlue.Print("  |   ")
	color.FgBlue.Print(spaces + "          " + highlightThing)
	color.FgRed.Print("      Error: ")
	fmt.Println("Pipe error")
}
