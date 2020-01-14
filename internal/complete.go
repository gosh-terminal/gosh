package internal

import (
	"bufio"
	"github.com/c-bata/go-prompt"
	"io/ioutil"
	"os"
	"strings"
)

// Unique remove prompt suggestion duplicates
func Unique(intSlice []prompt.Suggest) []prompt.Suggest {
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

// Completer complete commands
func Completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{{Text: "help", Description: "gosh"}, {Text: "exit", Description: "gosh"}, {Text: "history", Description: "gosh"}, {Text: "clearhist", Description: "gosh"}, {Text: "tree", Description: "gosh"}, {Text: "touch", Description: "gosh"}, {Text: "mkdir", Description: "gosh"}}
	var goshHome string = os.Getenv("GOSH_HOME")
	file, _ := os.Open(goshHome + "/history.txt")
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
		if file.IsDir() {
			s = append(s, prompt.Suggest{Text: file.Name(), Description: "Directory"})
		} else {
			s = append(s, prompt.Suggest{Text: file.Name(), Description: "File"})
		}
	}
	return prompt.FilterHasPrefix(Unique(s), d.GetWordBeforeCursor(), true)
}

// GetCommandHist Get command History
func GetCommandHist() []string {
	s := []string{}
	var goshHome string = os.Getenv("GOSH_HOME")
	file, _ := os.Open(goshHome + "/history.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Compare(string(scanner.Text()), "") == 0 {
			continue
		}
		s = append(s, scanner.Text())
	}
	return s
}
