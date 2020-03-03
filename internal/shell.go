package internal

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"github.com/mattn/go-isatty"
)


// Shell gosh shell
func Shell(quiet bool) {
	if !quiet {
		fmt.Println("Welcome to gosh, the Go Shell!")
		fmt.Println("------------------------------")
	}
	for {
		if !interactive() {
			fmt.Print(
				"\033[0;31m❌  gosh: This isn't an interactive terminal! Try `gosh -c COMMAND`.",
			)
		}
		ThePrompt()
		command := prompt.Input("", Completer, prompt.OptionHistory(GetCommandHist()), prompt.OptionSuggestionBGColor(prompt.DefaultColor),
			prompt.OptionInputTextColor(prompt.Cyan),
			prompt.OptionMaxSuggestion(6),
			prompt.OptionTitle("gosh"),
			prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
			prompt.OptionAddKeyBind(prompt.KeyBind{
				Key: prompt.ControlC,
				Fn: func(buf *prompt.Buffer) {
					ThePrompt()
				}}),
			prompt.OptionPreviewSuggestionTextColor(prompt.DefaultColor),
			prompt.OptionScrollbarBGColor(prompt.DefaultColor))
		command = strings.Replace(command, "\n", "", -1)
		if strings.Contains(command, "~") {
			command = strings.ReplaceAll(command, "~", os.Getenv("HOME"))
		} else if strings.Contains(command, "$") {
			regex, _ := regexp.Compile(".*\\$(.*).*")
			matches := regex.FindAllStringSubmatch(command, -1)
			for i, j := range matches {
				command1 := strings.Replace(command, j[i+1][0:], os.Getenv(j[i+1]), -1)
				command = strings.Replace(command1, "$", "", -1)
				Evaluate(command)
			}
		} else if strings.Contains(command, " > ") {
			data, err := SplitCommandFile(command)
			if err != nil {
				PipeError(command)
				UpdateHistory(command)
				continue
			}
			CaptureOutput, err := CaptureOutput(data[0])
			if err != nil {
				CommandNotFound(command)
			}
			RedirectToFile(CaptureOutput, data[1])
			UpdateHistory(command)
			continue
		} else if strings.Contains(command, " | ") {
			data, err := SplitCommands(command)
			if err != nil {
				PipeError(command)
				UpdateHistory(command)
				continue
			}
			CaptureOutput, err := CaptureOutput(data[0])
			if err != nil {
				CommandNotFound(command)
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
				DirectoryNotFound(dir)
			}
			err := os.Chdir(dir)
			if err != nil {
				if strings.HasSuffix(string(err.Error()), "file or directory") {
					DirectoryNotFound(dir)
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
					CommandNotFound(command)
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
			DirectoryNotFound(dir)
		}
		err := os.Chdir(dir)
		if err != nil {
			if strings.HasSuffix(string(err.Error()), "file or directory") {
				DirectoryNotFound(dir)
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
		data, err := SplitCommandFile(command)
		if err != nil {
			PipeError(command)
			UpdateHistory(command)
			return
		}
		CaptureOutput, err := CaptureOutput(data[0])
		if err != nil {
			CommandNotFound(command)
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
				CommandNotFound(command)
			}
		}
		UpdateHistory(command)
	}
}

func interactive() bool {
	return isatty.IsCygwinTerminal(os.Stdout.Fd()) || isatty.IsTerminal(os.Stdout.Fd())
}
