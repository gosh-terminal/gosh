package main
import (
    "fmt"
    "os"
    "strings"
    "bufio"
    "os/exec"
    "log"
)
func executeCommand(theCommand string) {
	cmd := exec.Command(theCommand)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Print(err)
	}
}
func getArg(commandString string) string {
	s:=strings.Split(commandString," ")
	if len(s) > 1 {
		return s[1]
	}
	return "error"
}
func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Welcome to gosh the Go Shell!")
    fmt.Println("-----------------------------")
    for {
        executeCommand("/workspace/gosh/src/commands/bin/prompt")
        command,_ := reader.ReadString('\n')
        command = strings.Replace(command, "\n", "", -1)
        if strings.Compare("help",command) == 0 {
			executeCommand("/workspace/gosh/src/commands/bin/help")
        } else if strings.Compare("exit",command) == 0 {
            break;
		} else if(strings.Compare("ls", command)==0) {
			executeCommand("/workspace/gosh/src/commands/list.py")
		} else if(strings.Compare("",command) == 0) {
			continue;
		} else if strings.HasPrefix(command,"cd") {
			dir := getArg(command)
			if dir == "error" {
				println("gosh: cd: file not specified")
			}
			os.Chdir(dir)
		} else if strings.HasPrefix(command,"history") {
			executeCommand("/workspace/gosh/src/commands/history.py")
			continue
		} else {
			cmd := exec.Command(command)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
            if err := cmd.Run(); err != nil {
                fmt.Println("gosh: " + command + ": command not found")
            }
		}
		f, err := os.OpenFile("/workspace/gosh/src/commands/history.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString("\n"+command); err != nil {
			log.Println(err)
		}
	}
}
