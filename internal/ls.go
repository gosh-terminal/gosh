package internal

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Ls this is the ls command for gosh
func Ls(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		DirectoryNotFound(path)
		return
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

type nodeValue string

func (nv nodeValue) String() string {
	return string(nv)
}

func getFiles(path string) []*widgets.TreeNode {
	nodes := []*widgets.TreeNode{}
	// Gets an array of files in the current directory
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		// Create a file item
		tmp := widgets.TreeNode{}
		tmp.Value = nodeValue(file.Name())
		// If its a directory recursivly recurse the subdirectories
		if file.IsDir() {
			tmp1 := getFiles(path + "/" + file.Name())
			tmp.Nodes = tmp1
		}
		// When the file item is done being processed add it to the file list
		nodes = append(nodes, &tmp)
	}
	return nodes
}

// TreeView tree view command
func TreeView(path string, tabNumbers int) {
	if err := ui.Init(); err != nil {
		log.Printf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	// This will hide the cursor
	fmt.Print("\033[?25l")
	nodes := getFiles(path)
	l := widgets.NewTree()
	l.TextStyle = ui.NewStyle(ui.ColorCyan)
	l.WrapText = false
	l.SetNodes(nodes)

	x, y := ui.TerminalDimensions()

	l.SetRect(0, 0, x, y)

	ui.Render(l)

	previousKey := ""
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "<Down>":
			l.ScrollDown()
		case "<Up>":
			l.ScrollUp()
		case "<Enter>":
			l.ToggleExpand()
		case "E":
			l.ExpandAll()
		case "C":
			l.CollapseAll()
		case "<Resize>":
			x, y := ui.TerminalDimensions()
			l.SetRect(0, 0, x, y)
		}

		if previousKey == "g" {
			previousKey = ""
		} else {
			previousKey = e.ID
		}

		ui.Render(l)
	}
}
