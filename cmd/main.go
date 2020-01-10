package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		shell()
	}
	if os.Args[1] == "-v" {
		fmt.Println("gosh v0.04-alpha")
	} else if os.Args[1] == "-c" {
		if len(os.Args) >= 3 {
			evaluate(os.Args[2])
		} else {
			invalidNumberOfArgs(os.Args[1])
		}
	}
}
