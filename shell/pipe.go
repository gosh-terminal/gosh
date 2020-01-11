package shell

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// RedirectToFile redorect the output of command to file
func RedirectToFile(commandOneOut []byte, fileName string) {
	err := ioutil.WriteFile(fileName, commandOneOut, 0655)
	if err != nil {
		os.Create(fileName)
		RedirectToFile(commandOneOut, fileName)
	}
}

// CaptureOutput capture the output of one command
func CaptureOutput(command string) ([]byte, error) {
	return exec.Command(command).Output()
}

// Pipe pipes output of one command to another
func Pipe(from []byte, to string) error {
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

// SplitCommands split pipes by ` | `
func SplitCommands(command string) ([]string, error) {
	splitStr := strings.Split(command, " | ")
	if len(splitStr) == 2 && splitStr[1] != "" && len(strings.TrimSpace(splitStr[1])) != 0 {
		return splitStr, nil
	}
	return splitStr, errors.New("Must specify a file")
}

// SplitCommandFile split the command and file by ` > `
func SplitCommandFile(command string) ([]string, error) {
	splitStr := strings.Split(command, " > ")
	if len(splitStr) == 2 && splitStr[1] != "" && len(strings.TrimSpace(splitStr[1])) != 0 {
		return splitStr, nil
	}
	return splitStr, errors.New("Must specify a file")
}
