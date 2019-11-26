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
	s := []prompt.Suggest{}
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
	for _, file := range files {
		s = append(s, prompt.Suggest{Text: file.Name()})
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
