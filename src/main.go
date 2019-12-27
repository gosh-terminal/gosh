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
		fmt.Println("gosh v0.02-alpha")
	} else if os.Args[1] == "-c" {
		evaluate(os.Args[2])
	}
}
