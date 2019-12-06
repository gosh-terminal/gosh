package main

import (
	"github.com/gookit/color"
	"os"
	"strings"
)

func mkdir(dirnames string) {
	dirname := strings.Split(dirnames, " ")
	for i := 1; i < len(dirname); i++ {
		os.Mkdir(dirname[i], 0655)
	}
	color.FgYellow.Println("the directories have been created! ✔")
}
