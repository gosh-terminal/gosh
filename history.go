package main
import (
	"os"
	"bufio"
	"strings"
	"fmt"
)
func history() {
	file, _ := os.Open("/workspace/gosh/data/history.txt")
	scanner := bufio.NewScanner(file)
	var num int = 1
	for scanner.Scan() {
		if strings.Compare(string(scanner.Text()), "") == 0 {
			continue
		}
		fmt.Printf("%d %s\n", num, scanner.Text())
		num++
	}
}