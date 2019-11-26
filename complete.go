package main

import (
	"bufio"
	"github.com/c-bata/go-prompt"
	"io/ioutil"
	"os"
	"strings"
)

func unique(intSlice []prompt.Suggest) []prompt.Suggest {
	keys := make(map[prompt.Suggest]bool)
	list := []prompt.Suggest{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{{Text: "help", Description: "Displays this help screen"}, {Text: "exit", Description: "Exit gosh"}, {Text: "history", Description: "displays commands you have previously run"}}
	var gopath string = os.Getenv("GOPATH")
	file, _ := os.Open(gopath + "/bin/history.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Compare(string(scanner.Text()), "") == 0 {
			continue
		}
		s = append(s, prompt.Suggest{Text: scanner.Text()})
	}
	files, _ := ioutil.ReadDir("/usr/bin")
	currentDir, _ := ioutil.ReadDir(".")
	for _, file := range files {
		s = append(s, prompt.Suggest{Text: file.Name(), Description: "Command"})
	}
	for _, file := range currentDir {
		s = append(s, prompt.Suggest{Text: file.Name(), Description: "File"})
	}
	return prompt.FilterHasPrefix(unique(s), d.GetWordBeforeCursor(), true)
}
func getCommandHist() []string {
	s := []string{}
	var gopath string = os.Getenv("GOPATH")
	file, _ := os.Open(gopath + "/bin/history.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Compare(string(scanner.Text()), "") == 0 {
			continue
		}
		s = append(s, scanner.Text())
	}
	return s
}
