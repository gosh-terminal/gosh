package main

import (
	"fmt"
	"github.com/gookit/color"
	"os"
	"strings"
)

func touch(filename string) {
	filenames := strings.Split(filename, " ")
	for i := 1; i < len(filenames); i++ {
		os.Create(filenames[i])
	}
	fmt.Printf("%s", color.FgYellow.Render("the files have been created! ✔\n"))
}
