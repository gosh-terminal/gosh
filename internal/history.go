package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ClearHistory clears the command history
func ClearHistory() string {
	var goshHome string = os.Getenv("GOSH_HOME")
	f, _ := os.OpenFile(goshHome+"/history.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.Truncate(0)
	fmt.Println("\033[0;32mHistory has been cleared ✔\033[0m")
	return "History has been cleared ✔\n"
}

// History the history command
func History() {
	var goshHome string = os.Getenv("GOSH_HOME")
	file, _ := os.Open(goshHome + "/history.txt")
	scanner := bufio.NewScanner(file)
	var num = 1
	fmt.Println("   \033[0;32m#      command\033[0m")
	fmt.Println(" ╭━━━━━━━━━━━━━━━━━━━╮")
	for scanner.Scan() {
		z := 14 - len(scanner.Text())
		spaces := strings.Repeat(" ", z)
		if string(scanner.Text()) == "" {
			continue
		}
		if num < 10 {
			fmt.Println(" │ \033[0;32m" + strconv.Itoa(num) + " \033[0m │ " + scanner.Text() + spaces + "│")
		} else {
			fmt.Println(" │ \033[0;32m" + strconv.Itoa(num) + " \033[0m│ " + scanner.Text() + spaces + "│")
		}
		num++
	}
	fmt.Println(" ╰━━━━━━━━━━━━━━━━━━━╯")
}

// UpdateHistory update the command history
func UpdateHistory(command string) {
	var goshHome string = os.Getenv("GOSH_HOME")
	f, err := os.OpenFile(goshHome+"/history.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString("\n" + command); err != nil {
		log.Println(err)
	}
}
