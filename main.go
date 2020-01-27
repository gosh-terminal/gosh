package main

import (
	"strings"
	"regexp"
	"io/ioutil"
	"fmt"
	shell "gosh/internal"
	"os"
)

func init() {
	if os.Getenv("GOSH_HOME") == "" {
		os.Setenv("GOSH_HOME", os.Getenv("HOME"))
	}
}

func main() {
	if len(os.Args) == 1 {
		shell.Shell()
	}
	if os.Args[1] == "-v" {
		fmt.Println("gosh v0.06-alpha")
	} else if os.Args[1] == "-c" {
		if len(os.Args) >= 3 {
			shell.Evaluate(os.Args[2])
		} else {
			shell.InvalidNumberOfArgs(os.Args[1])
		}
	} else if os.Args[1] == "run" {
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
	}
}
