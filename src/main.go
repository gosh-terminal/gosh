package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		shell()
	}
	switch os.Args[1] {
	case "-v":
		fmt.Println("gosh v0.02-alpha")
	case "-h":
		fmt.Println("gosh\n-v this will give the version")
	}
}
