package main
import (
    "fmt"
    "os"
    "strings"
    "bufio"
    "os/exec"
    "log"
)
func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Welcome to gosh the Go Shell!")
    fmt.Println("-----------------------------")
    for {
        fmt.Print("gosh> ")
        command,_ := reader.ReadString('\n')
        command = strings.Replace(command, "\n", "", -1)
        if strings.Compare("help",command) == 0 {
            cmd := exec.Command("nim c -r commands/help.nim")
            cmd.Stdout = os.Stdout
            cmd.Stderr = os.Stderr
            if err := cmd.Run(); err != nil {
                log.Fatal(err)
            }
        } else if strings.Compare("exit",command) == 0 {
            break;
		} else if(strings.Compare("listall", command)==0) {
            cmd := exec.Command("ls")
            cmd.Stdout = os.Stdout
            cmd.Stderr = os.Stderr
            if err := cmd.Run(); err != nil {
                log.Fatal(err)
            }
        }
    }
}