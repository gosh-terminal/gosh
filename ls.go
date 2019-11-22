package main

import (
	"fmt"
	"github.com/gookit/color"
	"io/ioutil"
	"log"
	"strings"
)

func ls() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	blue := color.FgCyan.Render
	yellow := color.FgYellow.Render
	fmt.Println("--------------------------------------------------")
	for _, file := range files {
		z := 25 - len(file.Name()) - -3
		spaces := strings.Repeat(" ", z)
		if file.IsDir() {
			fmt.Printf("| %s%s|| Directory       |\n", blue(file.Name()), spaces)
		} else if file.Mode().String() == "-rwxr-xr-x" {
			fmt.Printf("| %s%s|| Executable File |\n", yellow(file.Name()), spaces)
		} else {
			fmt.Printf("| %s%s|| File            |\n", file.Name(), spaces)
		}
	}
	fmt.Println("--------------------------------------------------")
}
