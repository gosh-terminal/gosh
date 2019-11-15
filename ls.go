package main
import (
	"log"
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
	yellow := color.FgYellow.Render
	fmt.Println("----------------------------------")
	for _, file := range files {
		z := 34 - len(file.Name()) - 3
		spaces := strings.Repeat(" ", z)
		if file.IsDir() {
			fmt.Printf("| %s%s|\n", blue(file.Name()), spaces)
		} else if file.Mode().String() == "-rwxr-xr-x" {
			fmt.Printf("| %s%s|\n", yellow(file.Name()), spaces)
		} else {
			fmt.Printf("| %s%s|\n", file.Name(), spaces)
		}
	}
	fmt.Println("----------------------------------")
}