package main

import (
	"fmt"
	shell "gosh/internal"
	"os"
)
func init()  {
	if os.Getenv("GOSH_HOME") == "" {
		os.Setenv("GOSH_HOME", os.Getenv("HOME"))
	}
}

func main() {
	if len(os.Args) == 1 {
		shell.Shell()
	}
	if os.Args[1] == "-v" {
		fmt.Println("gosh v0.06-alpha")
	} else if os.Args[1] == "-c" {
		if len(os.Args) >= 3 {
			shell.Evaluate(os.Args[2])
		} else {
			shell.InvalidNumberOfArgs(os.Args[1])
		}
	}
}
