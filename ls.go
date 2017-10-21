package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func serveDir(dir string) {
	f, err := os.OpenFile(dir, os.O_RDONLY, 0666)
	checkerr(err)
	files, err := f.Readdirnames(0)
	checkerr(err)
	for _, file := range files {
		fmt.Println(file)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "ls ( go implemenetation )"
	app.Usage = "ls <flags> <args>"
	app.Author = "Viswanath Kumar Skand Priya"
	app.Email = "kspviswa.github@gmail.com"
	app.Version = "0.1"
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "l, long"},
		cli.BoolFlag{Name: "d, dronly"},
	}
	app.Commands = []cli.Command{
		{
			Name:    "tree",
			Aliases: []string{"tr"},
			Usage:   "perform recursive directory lookup",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.Bool("long") {
			fmt.Println("Long value in flag")
		}
		if c.NArg() > 0 {
			serveDir(c.Args()[0])
		} else {
			serveDir(".")
		}
		return nil
	}

	app.Run(os.Args)
}
