package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func executeCommand(theCommand string) int {
	cmd := exec.Command(theCommand)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if err := cmd.Run() ; err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				return exitError.ExitCode()
			}
		}
	}
	return -2
}
func getArg(commandString string) string {
	var s []string = strings.Split(commandString, " ")
	if len(s) > 1 {
		return s[1]
	}
	return "error"
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to gosh the Go Shell!")
	fmt.Println("-----------------------------")
	for {
		executeCommand("/workspace/gosh/src/commands/bin/prompt")
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Could not read command!")
		}
		command = strings.Replace(command, "\n", "", -1)
		if strings.Compare("help", command) == 0 {
			executeCommand("/workspace/gosh/src/commands/bin/help")
		} else if strings.Compare("exit", command) == 0 {
			break
		} else if strings.Compare("ls", command) == 0 {
			executeCommand("/workspace/gosh/src/commands/list.py")
		} else if strings.Compare("", command) == 0 {
			continue
		} else if strings.HasPrefix(command, "cd") {
			var dir string = getArg(command)
			if dir == "error" {
				println("gosh: cd: directory not specified")
			}
			os.Chdir(dir)
		} else if strings.HasPrefix(command, "history") {
			executeCommand("/workspace/gosh/src/commands/history.py")
			continue
		} else {
			cmd := exec.Command(command)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Start(); err != nil {
				executeCommand("/workspace/gosh/src/commands/bin/makeRed")
				fmt.Println("gosh: " + command + ": command not found")
				executeCommand("/workspace/gosh/src/commands/bin/resetColor")
			}
		}
		f, err := os.OpenFile("/workspace/gosh/src/commands/history.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString("\n" + command); err != nil {
			log.Println(err)
		}
	}
}
