package shell

import (
	"fmt"
	"os"
	"strings"
)

// Touch Create File
func Touch(filename string) {
	filenames := strings.Split(filename, " ")
	if len(filenames) <= 1 {
		InvalidNumberOfArgs(filename)
		return
	}
	for i := 1; i < len(filenames); i++ {
		os.Create(filenames[i])
	}
	if len(filenames) >= 3 {
		fmt.Println("\033[0;33mThe files have been created! ✔\033[0m")
	} else {
		fmt.Println("\033[0;33m"+filenames[1], "has been created! ✔\033[0m")
	}
}
