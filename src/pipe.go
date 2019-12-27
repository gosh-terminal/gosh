package main

import (
	"bytes"
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

func pipe(from []byte, to string) error {
	cmd := exec.Command(to)

	buffer := bytes.Buffer{}
	buffer.Write(from)

	cmd.Stdout = os.Stdout
	cmd.Stdin = &buffer
	cmd.Stderr = os.Stderr

	err1 := cmd.Run()
	if err1 != nil {
		return err1
	}
	return nil
}

func splitCommands(command string) ([]string, error) {
	splitStr := strings.Split(command, " | ")
	if len(splitStr) == 2 && splitStr[1] != "" && len(strings.TrimSpace(splitStr[1])) != 0 {
		return splitStr, nil
	}
	return splitStr, errors.New("Must specify a file")
}

func splitCommandFile(command string) ([]string, error) {
	splitStr := strings.Split(command, " > ")
	if len(splitStr) == 2 && splitStr[1] != "" && len(strings.TrimSpace(splitStr[1])) != 0 {
		return splitStr, nil
	}
	return splitStr, errors.New("Must specify a file")
}
