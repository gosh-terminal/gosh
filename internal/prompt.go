package internal

import (
	"bufio"
	"fmt"
	"github.com/manifoldco/promptui"
	git "gopkg.in/src-d/go-git.v4"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// ThePrompt This function is for the main prompt of the shell
func ThePrompt() {
	isGitRepo := false
	r, err := git.PlainOpenWithOptions(".git", &git.PlainOpenOptions{DetectDotGit: true})
	if err == nil {
		isGitRepo = true
	}
	regex, _ := regexp.Compile("https://.*/.*/(.*)\\.git")
	list, err := r.Remotes()
	repoName := regex.FindStringSubmatch(list[0].String())
	dir, _ := os.Getwd()
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
	if isGitRepo {
		Rawhead, _ := r.Head()
		head1 := strings.Split(Rawhead.String(), "/")
		head := head1[len(head1)-1]
		fmt.Print("\033[0;36m(" + dir + ")\033[0;35m" + repoName[1] + "@\033[0;33m" + head + "\033[0;34 \033[032m gosh \033[0;34mλ \033[0m")
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
