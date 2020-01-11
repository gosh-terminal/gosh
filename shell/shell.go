package shell

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"os"
	"strings"
)

func Shell() {
	fmt.Println("Welcome to gosh, the Go Shell!")
	fmt.Println("------------------------------")
	for {
		ThePrompt()
		command := prompt.Input("", completer, prompt.OptionHistory(getCommandHist()), prompt.OptionSuggestionBGColor(prompt.DefaultColor),
			prompt.OptionInputTextColor(prompt.Cyan),
			prompt.OptionMaxSuggestion(6),
			prompt.OptionTitle("gosh"),
			prompt.OptionAddKeyBind(prompt.KeyBind{
				Key: prompt.ControlC,
				Fn: func(buf *prompt.Buffer) {
					ThePrompt()
				}}),
			prompt.OptionPreviewSuggestionTextColor(prompt.DefaultColor),
			prompt.OptionScrollbarBGColor(prompt.DefaultColor))
		command = strings.Replace(command, "\n", "", -1)
		if strings.Contains(command, " > ") {
			data, err := splitCommandFile(command)
			if err != nil {
				pipeError(command)
				UpdateHistory(command)
				continue
			}
			CaptureOutput, err := CaptureOutput(data[0])
			if err != nil {
				commandNotFound(command)
			}
			RedirectToFile(CaptureOutput, data[1])
			UpdateHistory(command)
			continue
		} else if strings.Contains(command, " | ") {
			data, err := splitCommands(command)
			if err != nil {
				pipeError(command)
				UpdateHistory(command)
				continue
			}
			CaptureOutput, err := CaptureOutput(data[0])
			if err != nil {
				commandNotFound(command)
			}
			Pipe(CaptureOutput, data[1])
			UpdateHistory(command)
			continue
		} else if strings.Compare("help", command) == 0 {
			Help()
			UpdateHistory(command)
		} else if strings.Compare("exit", command) == 0 {
			Exit()
		} else if strings.Compare("ls", command) == 0 {
			Ls(".")
			UpdateHistory(command)
		} else if strings.HasPrefix(command, "ls ") {
			var dir string = GetArg(command)
			if len(dir) == 0 {
				InvalidNumberOfArgs(command)
				continue
			}
			Ls(dir)
			UpdateHistory(command)
		} else if strings.Compare("", command) == 0 {
			continue
		} else if strings.HasPrefix(command, "cd ") {
			var dir string = GetArg(command)
			if dir == "error" {
				directoryNotFound(dir)
			}
			err := os.Chdir(dir)
			if err != nil {
				if strings.HasSuffix(string(err.Error()), "file or directory") {
					directoryNotFound(dir)
				}
			}
			UpdateHistory(command)
			continue
		} else if strings.HasPrefix("cd", command) {
			InvalidNumberOfArgs(command)
			continue
		} else if strings.HasPrefix(command, "history") {
			History()
			continue
		} else if strings.Compare(command, "clearhist") == 0 {
			ClearHistory()
		} else if command == "tree" {
			TreeView(".", 0)
			UpdateHistory(command)
			continue
		} else if strings.HasPrefix(command, "touch") {
			Touch(command)
			UpdateHistory(command)
			continue
		} else if strings.HasPrefix(command, "mkdir") {
			Mkdir(command)
			UpdateHistory(command)
			continue
		} else {
			if err := ExecuteCommand(command); err != nil {
				if strings.HasSuffix(string(err.Error()), "executable file not found in $PATH") {
					commandNotFound(command)
				}
			}
			UpdateHistory(command)
		}
	}
}
// Evaluate evaluate function
func Evaluate(command string) {
	if len(strings.Trim(command, " ")) == 0 {
		fmt.Println("Error input empty")
		return
	}
	if strings.Compare("help", command) == 0 {
		Help()
		UpdateHistory(command)
	} else if strings.Compare("exit", command) == 0 {
		Exit()
	} else if strings.Compare("ls", command) == 0 {
		Ls(".")
		UpdateHistory(command)
	} else if strings.HasPrefix(command, "ls ") {
		var dir string = GetArg(command)
		Ls(dir)
		UpdateHistory(command)
	} else if strings.Compare("", command) == 0 {
		return
	} else if strings.HasPrefix(command, "cd ") {
		var dir string = GetArg(command)
		if dir == "error" {
			directoryNotFound(dir)
		}
		err := os.Chdir(dir)
		if err != nil {
			if strings.HasSuffix(string(err.Error()), "file or directory") {
				directoryNotFound(dir)
			}
		}
		UpdateHistory(command)
		return
	} else if strings.HasPrefix(command, "history") {
		History()
		return
	} else if strings.Compare(command, "clearhist") == 0 {
		ClearHistory()
	} else if strings.Contains(command, " > ") {
		data, err := splitCommandFile(command)
		if err != nil {
			pipeError(command)
			UpdateHistory(command)
			return
		}
		CaptureOutput, err := CaptureOutput(data[0])
		if err != nil {
			commandNotFound(command)
		}
		RedirectToFile(CaptureOutput, data[1])
		UpdateHistory(command)
		return
	} else if command == "tree" {
		TreeView(".", 0)
		UpdateHistory(command)
		return
	} else if strings.HasPrefix(command, "touch") {
		Touch(command)
		UpdateHistory(command)
		return
	} else if strings.HasPrefix(command, "mkdir") {
		Mkdir(command)
		UpdateHistory(command)
		return
	} else {
		if err := ExecuteCommand(command); err != nil {
			if strings.HasSuffix(string(err.Error()), "executable file not found in $PATH") {
				commandNotFound(command)
			}
		}
		UpdateHistory(command)
	}
}
