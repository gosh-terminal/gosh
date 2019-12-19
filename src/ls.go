package main

import (
	"fmt"
	"github.com/gookit/color"
	"io/ioutil"
	"log"
	"strings"
)

func ls(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	blue := color.FgCyan.Render
	yellow := color.FgYellow.Render
	fmt.Printf("    %s     %s                             %s \n", color.FgGreen.Render("#"), color.FgGreen.Render("name"), color.FgGreen.Render("type"))
	fmt.Println(" ╭━━━━┯━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┯━━━━━━━━━━━━━━━━━╮")
	for i, file := range files {
		i++
		x := 3 - len(string(i))
		z := 33 - len(file.Name())
		spaces := strings.Repeat(" ", z)
		spaces2 := strings.Repeat(" ", x)
		if file.IsDir() {
			if i >= 10 {
				fmt.Printf(" │ %s%s│ %s%s│ Directory       │\n", color.FgGreen.Render(i), spaces2, blue(file.Name()), spaces)
			} else {
				fmt.Printf(" │ %s%s │ %s%s│ Directory       │\n", color.FgGreen.Render(i), spaces2, blue(file.Name()), spaces)
			}
		} else if file.Mode().String() == "-rwxr-xr-x" {
			if i >= 10 {
				fmt.Printf(" │ %s%s│ %s%s│ Executable File │\n", color.FgGreen.Render(i), spaces2, yellow(file.Name()), spaces)
			} else {
				fmt.Printf(" │ %s%s │ %s%s│ Executable File │\n", color.FgGreen.Render(i), spaces2, yellow(file.Name()), spaces)
			}
		} else {
			if i >= 10 {
				fmt.Printf(" │ %s%s│ %s%s│ File            │\n", color.FgGreen.Render(i), spaces2, file.Name(), spaces)
			} else {
				fmt.Printf(" │ %s%s │ %s%s│ File            │\n", color.FgGreen.Render(i), spaces2, file.Name(), spaces)
			}
		}
	}
	fmt.Println(" ╰━━━━━┷━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┷━━━━━━━━━━━━━━━━╯")
}

func treeView(path string, tabNumbers int) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("UNKNOWN ERROR HAS OCCURED!")
	}
	for _, file := range files {
		tabsThing := strings.Repeat("    ", tabNumbers)
		fmt.Println(tabsThing, file.Name())
		if file.IsDir() {
			treeView(path + "/" + file.Name(), tabNumbers+1)
		}
	}
}
