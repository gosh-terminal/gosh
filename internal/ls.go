package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Ls this is the ls command for gosh
func Ls(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("    \033[0;32m#     name                             type \033[0m")
	fmt.Println(" ╭━━━━┯━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┯━━━━━━━━━━━━━━━━━╮")
	for i, file := range files {
		i++
		x := 3 - len(string(i))
		z := 33 - len(file.Name())
		spaces := strings.Repeat(" ", z)
		spaces2 := strings.Repeat(" ", x)
		if file.IsDir() {
			if i >= 10 {
				fmt.Println(" │ \033[0;32m" + strconv.Itoa(i) + spaces2 + "\033[0m│ \033[0;36m" +
					file.Name() + spaces + "\033[0m│ Directory       \033[0m│")
			} else {
				fmt.Println(" │ \033[0;32m" + strconv.Itoa(i) + spaces2 + "\033[0m │ \033[0;36m" +
					file.Name() + spaces + "\033[0m│ Directory       \033[0m│")
			}
		} else if file.Mode().String() == "-rwxr-xr-x" {
			if i >= 10 {
				fmt.Println(" │ \033[0;32m" + strconv.Itoa(i) + spaces2 + "\033[0m│ \033[0;33m" +
					file.Name() + spaces + "\033[0m│ File            \033[0m│")
			} else {
				fmt.Println(" │ \033[0;32m" + strconv.Itoa(i) + spaces2 + "\033[0m │ \033[0;33m" +
					file.Name() + spaces + "\033[0m│ File            \033[0m│")
			}
		} else {
			if i >= 10 {
				fmt.Println(" │ \033[0;32m" + strconv.Itoa(i) + spaces2 + "\033[0m│ " + file.Name() + spaces +
					"\033[0m│ File            \033[0m│")
			} else {
				fmt.Println(" │ \033[0;32m" + strconv.Itoa(i) + spaces2 + "\033[0m │ " + file.Name() + spaces +
					"\033[0m│ File            \033[0m│")
			}
		}
	}
	fmt.Println(" ╰━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━╯")
}

// TreeView tree view command
func TreeView(path string, tabNumbers int) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("UNKNOWN ERROR HAS OCCURED!")
	}
	for _, file := range files {
		tabsThing := strings.Repeat("    ", tabNumbers)
		fmt.Println(tabsThing, file.Name())
		if file.IsDir() {
			TreeView(path+"/"+file.Name(), tabNumbers+1)
		}
	}
}
