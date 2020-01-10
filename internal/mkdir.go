package main

import (
	"fmt"
	"os"
	"strings"
)

func mkdir(dirnames string) {
	dirname := strings.Split(dirnames, " ")
	if len(dirname) <= 1 {
		invalidNumberOfArgs(dirnames)
		return
	}
	for i := 1; i < len(dirname); i++ {
		os.Mkdir(dirname[i], 0655)
	}
	if len(dirname) >= 3 {
		fmt.Println("\033[0;33mThe directories have been created! ✔\033[0m")
	} else {
		fmt.Println("\033[0;33mDirectory", dirname[1], "has been created! ✔\033[0m")
	}
}
