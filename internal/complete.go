package internal

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
)

var filePathCompleter = completer.FilePathCompleter{
	IgnoreCase: true,
}

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
	var gopath string = os.Getenv("GOSH_HOME")
	file, _ := os.Open(gopath + "/history.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Compare(string(scanner.Text()), "") == 0 {
			continue
		}
		s = append(s, prompt.Suggest{Text: scanner.Text(), Description: "History"})
	}
	files, _ := ioutil.ReadDir("/usr/bin")
	for _, file := range files {
		s = append(s, prompt.Suggest{Text: file.Name(), Description: "Command"})
	}
	completions := filePathCompleter.Complete(d)
	for i := range completions {
		completions[i].Description = "File"
	}
	for _, i := range prompt.FilterHasPrefix(Unique(s), d.GetWordBeforeCursor(), true) {
		completions = append(completions, i)
	}
	return completions
}

// GetCommandHist Get command History
func GetCommandHist() []string {
	s := []string{}
	var gopath string = os.Getenv("GOSH_HOME")
	file, _ := os.Open(gopath + "/history.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Compare(string(scanner.Text()), "") == 0 {
			continue
		}
		s = append(s, scanner.Text())
	}
	return s
}
