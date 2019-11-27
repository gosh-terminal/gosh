package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func clearHistory() {
	var gopath string = os.Getenv("GOPATH")
	f, _ := os.OpenFile(gopath+"/bin/history.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.Truncate(0)
}
func history() {
	var gopath string = os.Getenv("GOPATH")
	file, _ := os.Open(gopath + "/bin/history.txt")
	scanner := bufio.NewScanner(file)
	var num = 1
	for scanner.Scan() {
		if strings.Compare(string(scanner.Text()), "") == 0 {
			continue
		}
		fmt.Printf("%d %s\n", num, scanner.Text())
		num++
	}
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
