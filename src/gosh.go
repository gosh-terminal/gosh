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
        cmd := exec.Command("./src/commands/bin/prompt")
            cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
            if err := cmd.Run(); err != nil {
				log.Print(err)
            }
        command,_ := reader.ReadString('\n')
        command = strings.Replace(command, "\n", "", -1)
        if strings.Compare("help",command) == 0 {
			cmd := exec.Command("./src/commands/bin/help")
            cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
            if err := cmd.Run(); err != nil {
				log.Print(err)
            }
        } else if strings.Compare("exit",command) == 0 {
            break;
		} else if(strings.Compare("listall", command)==0) {
            cmd := exec.Command("ls")
            cmd.Stdout = os.Stdout
            cmd.Stderr = os.Stderr
            if err := cmd.Run(); err != nil {
                log.Print(err)
            }
		} else if(strings.Compare("",command) == 0) {
			continue;
		}else {
			cmd := exec.Command(command)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
            if err := cmd.Run(); err != nil {
                fmt.Println("gosh: " + command + ": command not found")
            }
		}
	}
}