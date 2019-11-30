package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		she ll()
	}
	switch os.Args[1] {
	case "-v":
        fmt.Println("v0.02-alpha")
	}
	case "-h":
		fmt.Println("Welcome to goshell!")
}