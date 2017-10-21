package main

import (
	"fmt"
	"os"
	"strconv"
	//"strconv"
	"time"

	"code.cloudfoundry.org/bytefmt"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

type flagdef struct {
	isLong   bool
	isDronly bool
	isFOnly  bool
	isHr     bool
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func retsize(nsize int64, hr bool) string {
	if hr {
		return bytefmt.ByteSize(uint64(nsize))
	}
	return fmt.Sprintf("%v", nsize)
}

func serveDecoratedDir(dir string, flags flagdef) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Size", "Perms", "At", "Dir"})

	f, err := os.OpenFile(dir, os.O_RDONLY, 0666)
	checkerr(err)
	finfo, err := f.Readdir(0)
	checkerr(err)
	for _, item := range finfo {
		if flags.isDronly && item.Mode().IsRegular() {
			continue
		}
		if flags.isFOnly && item.IsDir() {
			continue
		}
		data := []string{
			item.Name(),
			retsize(item.Size(), flags.isHr),
			item.Mode().String(),
			item.ModTime().Format("2006-01-02 15:04:05"),
			strconv.FormatBool(item.IsDir()),
		}
		table.Append(data)
	}
	table.Render()
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
	app.Name = "ls ( implemented in golang )"
	app.Usage = "ls [flags] [command][args]"
	app.Author = "Viswanath Kumar Skand Priya"
	app.Email = "kspviswa.github@gmail.com"
	app.Version = "0.1"
	app.Copyright = "MIT Licensed"
	app.Compiled = time.Now()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "l, long",
			Usage: "include extended information",
		},
		cli.BoolFlag{
			Name:  "d, dronly",
			Usage: "include only directories",
		},
		cli.BoolFlag{
			Name:  "f, fileonly",
			Usage: "include only regular files",
		},
		cli.BoolFlag{
			Name:  "hr",
			Usage: "view size in humar readable format",
		},
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
		decor := c.Bool("long") || c.Bool("dronly") || c.Bool("fileonly") || c.Bool("hr")
		flags := flagdef{c.Bool("long"), c.Bool("dronly"), c.Bool("fileonly"), c.Bool("hr")}
		dir := "."
		if c.NArg() > 0 {
			dir = c.Args()[0]
		}
		if decor {
			serveDecoratedDir(dir, flags)
		} else {
			serveDir(dir)
		}
		return nil
	}

	app.Run(os.Args)
}
