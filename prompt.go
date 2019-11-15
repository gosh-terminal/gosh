package main
import "github.com/gookit/color"
func prompt() {
	blue := color.FgBlue.Render
	color.FgGreen.Printf("gosh %s ", blue("λ"))
}