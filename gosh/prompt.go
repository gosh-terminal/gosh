package main

import (
	"bufio"
	"fmt"
	"github.com/gookit/color"
	"io/ioutil"
	"os"
	"strings"
)

func thePrompt() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("ERROR")
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".goshc") {
			file1, _ := os.Open(file.Name())
			scanner := bufio.NewScanner(file1)
			for scanner.Scan() {
				if strings.HasPrefix(scanner.Text(), "prompt ") {
					testArray := strings.Split(scanner.Text(), "prompt ")
					fmt.Print(testArray[1])
					return
				}
			}
		}
	}
	blue := color.FgBlue.Render
	color.FgGreen.Printf("gosh %s ", blue("λ"))
}
