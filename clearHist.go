package main
import "os"
func clearHistory() {
	f, _ := os.OpenFile("/workspace/gosh/data/history.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.Truncate(0)
}