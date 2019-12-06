package main

import (
	"fmt"
	"github.com/gookit/color"
)

func help() {
	fmt.Printf("   %s   %s        %s\n", color.FgGreen.Render("#"), color.FgGreen.Render("command"), color.FgGreen.Render("description"))
	fmt.Println(" ╭━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╮")
	fmt.Printf(" │ %s │ help       │  displays this help msg                       │\n", color.FgGreen.Render("1"))
	fmt.Printf(" │ %s │ exit       │  exits the terminal                           │\n", color.FgGreen.Render("2"))
	fmt.Printf(" │ %s │ history    │  displays commands you have previously run    │\n", color.FgGreen.Render("3"))
	fmt.Printf(" │ %s │ clearhist  │  clears your command history                  │\n", color.FgGreen.Render("4"))
	fmt.Printf(" │ %s │ tree       │  shows files as tree view                     │\n", color.FgGreen.Render("5"))
	fmt.Printf(" │ %s │ touch      │  creates an new file                          │\n", color.FgGreen.Render("6"))
	fmt.Println(" ╰━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━╯")
}
