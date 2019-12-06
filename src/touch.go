package main

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

func touch(filename string) {
	os.Create(filename)
	fmt.Printf("%s", color.FgYellow.Render(filename, " has been created! ✔\n"))
}
