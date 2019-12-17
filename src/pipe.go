package main

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func redirectToFile(commandOneOut []byte, fileName string) {
	err := ioutil.WriteFile(fileName, commandOneOut, 0655)
	if err != nil {
		os.Create(fileName)
		redirectToFile(commandOneOut, fileName)
	}
}

func captureOutput(command string) ([]byte, error) {
	return exec.Command(command).Output()
}

func splitCommands(command string) []string {
	return strings.Split(command, " | ")
}

func pipeToOtherCommand() {
	// WIP
}

func splitCommandFile(command string) ([]string, error) {
	splitStr := strings.Split(command, " > ")
	if len(splitStr) == 2 && splitStr[1] != "" && len(strings.TrimSpace(splitStr[1])) != 0 {
		return splitStr, nil
	}
	return splitStr, errors.New("Must specify a file")
}
