package main

import (
	"bufio"
	"github.com/c-bata/go-prompt"
	"io/ioutil"
	"os"
	"strings"
)

// This function takes a suggestion slice turns it into a map removing any duplicate keys and
// then turns itb back into a slice
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

// This generates the autocompletions first gosh specific commands then the command history
// Then /usr/bin and finally the current directory
// Pre-conditions this assumes the user is using a linux system because of the use of env variables and /usr/bin
func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{{Text: "help", Description: "gosh"}, {Text: "exit", Description: "gosh"}, {Text: "history", Description: "gosh"}, {Text: "clearhist", Description: "gosh"}}
	var gopath string = os.Getenv("GOPATH")
	file, _ := os.Open(gopath + "/bin/history.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Compare(string(scanner.Text()), "") == 0 {
			continue
		}
		s = append(s, prompt.Suggest{Text: scanner.Text(), Description: "History"})
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

// returns a list of all items in the history
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
