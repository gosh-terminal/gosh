package main

import (
	"fmt"
	"github.com/gookit/color"
	"os"
	"strings"
)

func touch(filename string) {
	filenames := strings.Split(filename, " ")
	if len(filenames) <= 1 {
		invalidNumberOfArgs(filename)
		return
	}
	for i := 1; i < len(filenames); i++ {
		os.Create(filenames[i])
	}
	if len(filenames) >= 3 {
		fmt.Printf("%s", color.FgYellow.Render("the files have been created! ✔\n"))
	} else {
		color.FgYellow.Println("file", filenames[1], "has been created! ✔")
	}
}
