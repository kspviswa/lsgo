package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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
	isTree   bool
}

type fileRef struct {
	path string
	name string
}

var walk []fileRef

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

func inspect(spath string, finfo os.FileInfo, err error) error {
	var sname string
	if err != nil {
		sname = err.Error()
	}
	sname = finfo.Name()
	foo := fileRef{path: spath, name: sname}
	walk = append(walk, foo)
	return nil
}

func serveDecoratedDirTree(dir string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Path", "Name"})
	dirname := dir

	err := filepath.Walk(dir, inspect)
	checkerr(err)

	for _, item := range walk {
		dirname = item.path
		data := []string{
			filepath.Dir(dirname),
			item.name,
		}
		table.Append(data)
	}
	table.SetAutoMergeCells(true)
	table.SetAutoWrapText(true)
	table.SetRowLine(true)
	table.Render()
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
	app.Author = "Viswa Kumar"
	app.Email = ""
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
			Name:  "hr, humanfriendly",
			Usage: "view size in humar readable format",
		},
		cli.BoolFlag{
			Name:  "t, tree",
			Usage: "view tree structure ( recursive lookup )",
		},
	}
	app.Action = func(c *cli.Context) error {
		decor := c.Bool("long") || c.Bool("dronly") || c.Bool("fileonly") || c.Bool("hr") || c.Bool("t")
		flags := flagdef{c.Bool("long"), c.Bool("dronly"), c.Bool("fileonly"), c.Bool("hr"), c.Bool("t")}
		dir := "."
		if c.NArg() > 0 {
			dir = c.Args()[0]
		}
		if decor {
			if flags.isTree {
				serveDecoratedDirTree(dir)
			} else {
				serveDecoratedDir(dir, flags)
			}
		} else {
			serveDir(dir)
		}
		return nil
	}

	app.Run(os.Args)
}
