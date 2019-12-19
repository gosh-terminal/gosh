package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(c.App.Version)
	}
	app.Name = "gosh"
	app.Usage = "Do anything from the terminal!"
	app.Action = func(c *cli.Context) error {
		shell()
		return nil
	}
	app.Version = "gosh v0.02-alpha"
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
