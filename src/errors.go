package main

import (
  "fmt"
  "strings"
)

func invalidNumberOfArgs(command string) {
  stringLen := len(command)
  highlightThing := strings.Repeat("^", stringLen)
  fmt.Print("\033[0;34m1 | ")
  fmt.Println("\033[0;31m❌  gosh: " + command + "\033[0m")
  fmt.Print("\033[0;34m  |   ")
  fmt.Print("\033[1;34m       " + highlightThing + "\033[0m")
  fmt.Print("\033[0;31m      Error: \033[0m")
  fmt.Println("Invalid number of arguments!")
}

func commandNotFound(command string) {
  stringLen := len(command)
  highlightThing := strings.Repeat("^", stringLen)
  fmt.Print("\033[0;34m1 | ")
  fmt.Println("\033[0;31m❌  gosh: " + command + "\033[0m")
  fmt.Print("\033[0;34m  |   ")
  fmt.Print("\033[1;34m       " + highlightThing + "\033[0m")
  fmt.Print("\033[0;31m      Error: \033[0m")
  fmt.Println("Command not found!")
}

func directoryNotFound(dir string) {
  fmt.Println("\033[0;31m❌  gosh: '" + dir + "' not found!\033[0m")
}

func pipeError(command string) {
  indexOf := strings.Index(command, " > ")
  highlightThing := strings.Repeat("^", len(command[indexOf+3:])+1)
  spaces := strings.Repeat(" ", len(command[:indexOf])+1)
  fmt.Print("\033[0;34m1 | ")
  fmt.Println("\033[0;31m❌  gosh: " + " " + command + "\033[0m")
  fmt.Print("\033[0;34m  |   ")
  fmt.Print(spaces + " \033[1;34m          " + highlightThing + "\033[0m")
  fmt.Print("\033[0;31m      Error: \033[0m")
  fmt.Println("Pipe error!")
}
