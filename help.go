package main

import "github.com/gookit/color"

func help() {
	color.FgMagenta.Println("Commands:")
	color.FgYellow.Println("help: displays this help screen\nexit: exits the terminal\nhistory: displays commands you have previously run\nclearhist: clears your command history")
}
