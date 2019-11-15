package main

import (
	"bufio"
	"fmt"
	"github.com/gookit/color"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to gosh the Go Shell!")
	fmt.Println("-----------------------------")
	for {
		prompt()
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Could not read command!")
		}
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
		} else if strings.HasPrefix(command, "cd") {
			var dir string = getArg(command)
			if dir == "error" {
				color.FgRed.Println("gosh: cd: directory not specified")
			}
			os.Chdir(dir)
			updateHistory(command)
		} else if strings.HasPrefix(command, "history") {
			history()
			continue
		} else if strings.Compare(command, "clearhist") == 0 {
			clearHistory()
			fmt.Println("history has been cleared")
		} else {
			if err = executeCommand(command); err != nil {
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
