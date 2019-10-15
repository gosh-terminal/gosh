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
    fmt.Println("-----------------------------")
    for {
        fmt.Print(">>> ")
        command,_ := reader.ReadString('\n')
        command = strings.Replace(command, "\n", "", -1)
        if strings.Compare("help",command) == 0 {
            fmt.Println("Commands:\nhelp: displays this help screen\nexit:exits the terminal")
        } else if strings.Compare("exit",command) == 0 {
			break;
		}
    }
}