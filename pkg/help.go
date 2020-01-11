package pkg

import (
	"fmt"
)

// Help the help function
func Help() {
	fmt.Println("   \033[0;32m#   command        description\033[0m")
	fmt.Println(" ╭━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╮")
	fmt.Println(" │ \033[0;32m1\033[0m │ exit       │  exits gosh                                  │")
	fmt.Println(" │ \033[0;32m2\033[0m │ history    │  displays commands you have previously run   │")
	fmt.Println(" │ \033[0;32m3\033[0m │ clearhist  │  clears your command history                 │")
	fmt.Println(" │ \033[0;32m4\033[0m │ tree       │  shows files as tree view                    │")
	fmt.Println(" │ \033[0;32m5\033[0m │ touch      │  creates an new file                         │")
	fmt.Println(" │ \033[0;32m6\033[0m │ mkdir      │  creates an new directory                    │")
	fmt.Println(" ╰━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╯")
}
