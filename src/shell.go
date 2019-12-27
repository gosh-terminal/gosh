package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"os"
	"strings"
)

func shell() {
	fmt.Println("Welcome to gosh, the Go Shell!")
	fmt.Println("------------------------------")
	for {
		thePrompt()
		command := prompt.Input("", completer, prompt.OptionHistory(getCommandHist()), prompt.OptionSuggestionBGColor(prompt.DefaultColor),
			prompt.OptionInputTextColor(prompt.Cyan),
			prompt.OptionMaxSuggestion(6),
			prompt.OptionTitle("gosh"),
			prompt.OptionAddKeyBind(prompt.KeyBind{
				Key: prompt.ControlC,
				Fn: func(buf *prompt.Buffer) {
					thePrompt()
				}}),
			prompt.OptionPreviewSuggestionTextColor(prompt.DefaultColor),
			prompt.OptionScrollbarBGColor(prompt.DefaultColor))
		command = strings.Replace(command, "\n", "", -1)
		if strings.Compare("help", command) == 0 {
			help()
			updateHistory(command)
		} else if strings.Compare("exit", command) == 0 {
			exit()
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
		} else if strings.HasPrefix(command, "touch") {
			touch(command)
			updateHistory(command)
			continue
		} else if strings.HasPrefix(command, "mkdir") {
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
func evaluate(command string) {
	if strings.Compare("help", command) == 0 {
		help()
		updateHistory(command)
	} else if strings.Compare("exit", command) == 0 {
		exit()
	} else if strings.Compare("ls", command) == 0 {
		ls(".")
		updateHistory(command)
	} else if strings.HasPrefix(command, "ls ") {
		var dir string = getArg(command)
		ls(dir)
		updateHistory(command)
	} else if strings.Compare("", command) == 0 {
		return
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
		return
	} else if strings.HasPrefix(command, "history") {
		history()
		return
	} else if strings.Compare(command, "clearhist") == 0 {
		clearHistory()
	} else if strings.Contains(command, " > ") {
		data, err := splitCommandFile(command)
		if err != nil {
			pipeError(command)
			updateHistory(command)
			return
		}
		captureOutput, err := captureOutput(data[0])
		if err != nil {
			commandNotFound(command)
		}
		redirectToFile(captureOutput, data[1])
		updateHistory(command)
		return
	} else if command == "tree" {
		treeView(".", 0)
		updateHistory(command)
		return
	} else if strings.HasPrefix(command, "touch") {
		touch(command)
		updateHistory(command)
		return
	} else if strings.HasPrefix(command, "mkdir") {
		mkdir(command)
		updateHistory(command)
		return
	} else {
		if err := executeCommand(command); err != nil {
			if strings.HasSuffix(string(err.Error()), "executable file not found in $PATH") {
				commandNotFound(command)
			}
		}
		updateHistory(command)
	}
}