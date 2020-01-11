package internal

import (
	"os"
	"os/exec"
	"strings"
)

// ExecuteCommand executes a command string
func ExecuteCommand(theCommand string) error {
	args := strings.Split(theCommand, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

// GetArg Get argument from commandString
func GetArg(commandString string) string {
	var s []string = strings.Split(commandString, " ")
	if len(s) >= 1 {
		return s[1]
	}
	return "error"
}
