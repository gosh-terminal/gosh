package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	shell "gosh/internal"
	util "gosh/pkg"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func init() {
	if os.Getenv("GOSH_HOME") == "" {
		os.Setenv("GOSH_HOME", os.Getenv("HOME"))
	}
}

func main() {
	quiet := false

	if util.ItemExists(os.Args, "-q") || util.ItemExists(os.Args, "--quiet") {
		quiet = true
	}

	if util.ItemExists(os.Args, "-v") {
		fmt.Println("gosh v0.06-alpha")
	} else if util.ItemExists(os.Args, "-c") {
		if len(os.Args) >= 3 {
			shell.Evaluate(os.Args[2])
		} else {
			shell.InvalidNumberOfArgs(os.Args[1])
		}
	} else if util.ItemExists(os.Args, "run") {
		f, _ := ioutil.ReadFile(os.Args[2])
		f1 := string(f)
		f1 = strings.ReplaceAll(f1, "\n", "")
		content := strings.Split(string(f1), ";")
		for _, i := range content {
			isMatch, _ := regexp.MatchString(" *", i)
			if len(i) == 0 || !isMatch {
				continue
			}
			shell.Evaluate(i)
		}
	} else {
		if terminal.IsTerminal(int(os.Stdin.Fd())) {
			shell.Shell(quiet)
		} else {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			shell.Evaluate(strings.Trim(text, "\n"))
		}
	}
}
