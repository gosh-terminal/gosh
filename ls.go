package main
import (
	"log"
	"os"
	"io/ioutil"
	"fmt"
	"github.com/gookit/color"
	"strings"
)
func ls() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	blue := color.FgCyan.Render
	red := color.FgRed.Render
	green := color.FgGreen.Render
	white := color.FgWhite.Render
	red1 := color.FgRed.Render
	// yellow := color.FgYellow.Render
	fmt.Println("----------------------------------")
	for _, file := range files {
		z := 34 - len(file.Name()) - 3
		spaces := strings.Repeat(" ", z)
		if file.IsDir() {
			if os.Getenv("exFiles") == "blue" {
				
			} 
			fmt.Printf("| %s%s|\n", blue(file.Name()), spaces)
		} else if file.Mode().String() == "-rwxr-xr-x" {
			fmt.Printf("| %s%s|\n", style(file.Name()), spaces)
		} else {
			fmt.Printf("| %s%s|\n", file.Name(), spaces)
		}
	}
	fmt.Println("----------------------------------")
}

func setSyntax() {
	var input string
	fmt.Println("Welcome to syntax setter.  Currently, executable files are yellow, directories are blue, and files are white.  Possible colors are:")
	fmt.Println("Blue\nWhite\nRed\nGreen\nYellow\nSet executable file: ")
	fmt.Scan(&input)
	if strings.ToLower(input) == "blue" {
		os.Setenv("exFiles", "blue")
	} else if strings.ToLower(input) == "green" {
		os.Setenv("exFiles", "green")
	} else if strings.ToLower(input) == "red" {
		os.Setenv("exFiles", "red")
	} else if strings.ToLower(input) == "white" {
		os.Setenv("exFiles", "white")
	} else {
		fmt.Println("color not found")
	}
	fmt.Println("Set directories: ")
	fmt.Scan(&input)
	if strings.ToLower(input) == "blue" {
		os.Setenv("dirs", "blue")
	} else if strings.ToLower(input) == "green" {
		os.Setenv("dirs", "green")
	} else if strings.ToLower(input) == "red" {
		os.Setenv("dirs", "red")
	} else if strings.ToLower(input) == "white" {
		os.Setenv("dirs", "white")
	} else {
		fmt.Println("color not found")
	}
	fmt.Println("Set files: ")
	fmt.Scan(&input)
	if strings.ToLower(input) == "blue" {
		os.Setenv("file", "blue")
	} else if strings.ToLower(input) == "green" {
		os.Setenv("file", "green")
	} else if strings.ToLower(input) == "red" {
		os.Setenv("file", "red")
	} else if strings.ToLower(input) == "white" {
		os.Setenv("file", "white")
	} else {
		fmt.Println("color not found")
	}
}
