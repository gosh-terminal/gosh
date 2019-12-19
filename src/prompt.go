package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gookit/color"
	"github.com/manifoldco/promptui"
)

func thePrompt() {
	isGitRepo := false
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
		if file.IsDir() && file.Name() == ".git" {
			isGitRepo = true
		}
	}
	blue := color.FgBlue.Render
	if isGitRepo {
		file1, _ := os.Open(".git/HEAD")
		scanner := bufio.NewScanner(file1)
		var gitHeadString string
		for scanner.Scan() {
			gitHeadString += scanner.Text()
		}
		dataStr := strings.Split(strings.Trim(gitHeadString, "\n"), "/")
		color.FgGreen.Printf("gosh%s%s %s ", color.FgMagenta.Render("@"), color.FgYellow.Render(dataStr[len(dataStr)-1]), blue("λ"))
		return
	}
	color.FgGreen.Printf("gosh %s ", blue("λ"))
}

func exit() {
	prompt := promptui.Select{
		Label: "Are you sure?",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		println("Error, exiting...")
		os.Exit(1)
	}
	if result == "Yes" {
		os.Exit(0)
	} else {
		return
	}
}
