package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/gookit/color"
	"os"
	"strings"
)
func main() {
	fmt.Println("Welcome to gosh the Go Shell!")
	fmt.Println("-----------------------------")
	for {
		thePrompt()
		command := prompt.Input("", completer, prompt.OptionHistory(getCommandHist()), prompt.OptionSuggestionBGColor(prompt.DefaultColor),
			prompt.OptionInputTextColor(prompt.Cyan),
			prompt.OptionMaxSuggestion(4),
			prompt.OptionTitle("gosh"),
			prompt.OptionAddKeyBind(prompt.KeyBind{
			Key: prompt.ControlC,
			Fn: func(buf *prompt.Buffer) {
				os.Exit(0)
			}}),
			prompt.OptionPreviewSuggestionTextColor(prompt.DefaultColor),
			prompt.OptionScrollbarBGColor(prompt.DefaultColor))
		command = strings.Replace(command, "\n", "", -1)
		if strings.Compare("help", command) == 0 {
			help()
			updateHistory(command)
		} else if strings.Compare("exit", command) == 0 {
			os.Exit(1)
		} else if strings.Compare("ls", command) == 0 {
			ls()
			updateHistory(command)
		} else if strings.Compare("", command) == 0 {
			continue
		} else if strings.HasPrefix(command, "cd ") {
			var dir string = getArg(command)
			if dir == "error" {
				color.FgRed.Println("gosh: cd: directory not specified")
			}
			err := os.Chdir(dir)
			if err != nil {
				if strings.HasSuffix(string(err.Error()), "file or directory") {
					color.FgRed.Println("gosh: " + command + ": directory not found")
				}
			}
			updateHistory(command)
			continue
		} else if strings.HasPrefix(command, "history") {
			history()
			continue
		} else if strings.Compare(command, "clearhist") == 0 {
			clearHistory()
			fmt.Println("history has been cleared")
		} else if command == "setlscolor" {
			// WIP
		} else {
			if err := executeCommand(command); err != nil {
				if strings.HasSuffix(string(err.Error()), "executable file not found in $PATH") {
					color.FgRed.Println("gosh: " + command + ": command not found")
				}
			}
			updateHistory(command)
		}
	}
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("%s\n", err.Error()))
	}
}
