package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/manifoldco/promptui"
	"os"
	"strings"
)

func shell() {
	fmt.Println("Welcome to gosh the Go Shell!")
	fmt.Println("-----------------------------")
	for {
		thePrompt()
		command := prompt.Input("", completer, prompt.OptionHistory(getCommandHist()), prompt.OptionSuggestionBGColor(prompt.DefaultColor),
			prompt.OptionInputTextColor(prompt.Cyan),
			prompt.OptionMaxSuggestion(6),
			prompt.OptionTitle("gosh"),
			prompt.OptionAddKeyBind(prompt.KeyBind{
				Key: prompt.ControlC,
				Fn: func(buf *prompt.Buffer) {
					return
				}}),
			prompt.OptionPreviewSuggestionTextColor(prompt.DefaultColor),
			prompt.OptionScrollbarBGColor(prompt.DefaultColor))
		command = strings.Replace(command, "\n", "", -1)
		if strings.Compare("help", command) == 0 {
			help()
			updateHistory(command)
		} else if strings.Compare("exit", command) == 0 {
			prompt := promptui.Select{
				Label: "Are you sure?",
				Items: []string{"Yes", "No"},
			}
			_, result, err := prompt.Run()
			if err != nil {
				println("ERROR")
			}
			if result == "Yes" {
				os.Exit(0)
			} else {
				continue
			}
		} else if strings.Compare("ls", command) == 0 {
			ls(".")
			updateHistory(command)
		} else if strings.HasPrefix(command, "ls ") {
			var dir string = getArg(command)
			ls(dir)
			updateHistory(command)
		} else if strings.Compare("", command) == 0 {
			continue
		} else if strings.HasPrefix(command, "cd ") {
			var dir string = getArg(command)
			if dir == "error" {
				directoryNotFound(dir)
			}
			err := os.Chdir(dir)
			if err != nil {
				if strings.HasSuffix(string(err.Error()), "file or directory") {
					directoryNotFound(dir)
				}
			}
			updateHistory(command)
			continue
		} else if strings.HasPrefix(command, "history") {
			history()
			continue
		} else if strings.Compare(command, "clearhist") == 0 {
			clearHistory()
		} else if command == "setlscolor" {
			// WIP
		} else if strings.Contains(command, " > ") {
			data, err := splitCommandFile(command)
			if err != nil {
				pipeError(command)
				updateHistory(command)
				continue
			}
			captureOutput, err := captureOutput(data[0])
			if err != nil {
				commandNotFound(command)
			}
			redirectToFile(captureOutput, data[1])
			updateHistory(command)
			continue
		} else if command == "tree" {
			treeView(".", 0)
			updateHistory(command)
			continue
		} else if strings.HasPrefix(command, "touch ") {
			touch(command)
			updateHistory(command)
			continue
		} else if strings.HasPrefix(command, "mkdir ") {
			mkdir(command)
			updateHistory(command)
			continue
		} else {
			if err := executeCommand(command); err != nil {
				if strings.HasSuffix(string(err.Error()), "executable file not found in $PATH") {
					commandNotFound(command)
				}
			}
			updateHistory(command)
		}
	}
}
