package main
import  (
	"os"
	"os/exec"
	"strings"
)
func executeCommand(theCommand string) error {
	args := strings.Split(theCommand, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
func getArg(commandString string) string {
	var s []string = strings.Split(commandString, " ")
	if len(s) >= 1 {
		return s[1]
	}
	return "error"
}