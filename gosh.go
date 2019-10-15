package main
import (
    "fmt"
    "os"
    "strings"
    "bufio"
)
func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Welcome to gosh the Go Shell!")
    fmt.Println("--------------------------")
    for {
        fmt.Print(">>> ")
        text,_ := reader.ReadString('\n')
        text = strings.Replace(text, "\n", "", -1)
        if strings.Compare("help",text) == 0 {
            fmt.Println("Commands:\nhelp: displays this help screen")
        }
    }
}