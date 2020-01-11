package pkg

import (
	"bufio"
	"fmt"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"os"
	"strings"
)

// ThePrompt This function is for the main prompt of the shell
func ThePrompt() {
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
	if isGitRepo {
		file1, _ := os.Open(".git/HEAD")
		scanner := bufio.NewScanner(file1)
		var gitHeadString string
		for scanner.Scan() {
			gitHeadString += scanner.Text()
		}
		dataStr := strings.Split(strings.Trim(gitHeadString, "\n"), "/")
		fmt.Print("\033[0;32mgosh\033[0;35m@\033[0;33m" + dataStr[len(dataStr)-1] + "\033[0;34m λ \033[0m")
		return
	}
	fmt.Printf("\033[0;32mgosh \033[0;34mλ \033[0m")
}

// Exit This function exits the program
func Exit() {
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
