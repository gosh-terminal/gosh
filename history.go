package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gookit/color"
)

func clearHistory() {
	var gopath string = os.Getenv("GOPATH")
	f, _ := os.OpenFile(gopath+"/bin/history.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.Truncate(0)
	fmt.Printf("%s ✔\n", color.FgYellow.Render("History has been cleared"))
}
func history() {
	var gopath string = os.Getenv("GOPATH")
	file, _ := os.Open(gopath + "/bin/history.txt")
	scanner := bufio.NewScanner(file)
	var num = 1
	fmt.Printf("   %s      %s\n", color.FgGreen.Render("#"), color.FgGreen.Render("command"))
	fmt.Println(" ╭━━━━━━━━━━━━━━━━━━━╮")
	for scanner.Scan() {
		z := 14 - len(scanner.Text())
		spaces := strings.Repeat(" ", z)
		if strings.Compare(string(scanner.Text()), "") == 0 {
			continue
		}
		if num < 10 {
			fmt.Printf(" │ %s  │ %s%s│\n", color.FgGreen.Render(num), scanner.Text(), spaces)
		} else {
			fmt.Printf(" │ %s │ %s%s│\n", color.FgGreen.Render(num), scanner.Text(), spaces)
		}
		num++
	}
	fmt.Println(" ╰━━━━━━━━━━━━━━━━━━━╯")
}

func updateHistory(command string) {
	var gopath string = os.Getenv("GOPATH")
	f, err := os.OpenFile(gopath+"/bin/history.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString("\n" + command); err != nil {
		log.Println(err)
	}
}
