package main

import (
	"github.com/gookit/color"
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
		color.FgYellow.Println("the directories have been created! ✔")
	} else {
		color.FgYellow.Println("directory", dirname[1], "has been created! ✔")
	}
} 
