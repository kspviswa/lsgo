package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"code.cloudfoundry.org/bytefmt"
	"github.com/olekukonko/tablewriter"
	flag "github.com/spf13/pflag"
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

func returnSizeStr(isB bool) string {
	if isB {
		return "Size"
	}
	return "Size (Bytes)"
}

func serveDecoratedDir(dir string, flags flagdef) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", returnSizeStr(flags.isHr), "Perms", "At", "Dir"})

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

	var longB, drB, fB, hrB, tB bool

	flag.BoolVarP(&longB, "long", "l", false, "include extended information")
	flag.BoolVarP(&drB, "dronly", "d", false, "include only directories")
	flag.BoolVarP(&fB, "fileonly", "f", false, "include only regular files")
	flag.BoolVarP(&longB, "humanfriendly", "r", false, "view size in human readableformat")
	flag.BoolVarP(&longB, "tree", "t", false, "view tree structure ( recursive lookup ) ")

	// Using the posix compliant p-flag version

	/*
		// long format
		flag.BoolVar(&longB, "long", false, "include extended information")
		flag.BoolVar(&longB, "l", false, "include extended information")

		// Drive only format
		flag.BoolVar(&drB, "d", false, "include only directories")
		flag.BoolVar(&drB, "dronly", false, "include only directories")

		// File only format
		flag.BoolVar(&fB, "f", false, "include only regular files")
		flag.BoolVar(&fB, "fileonly", false, "include only regular files")

		// Human readable format
		flag.BoolVar(&hrB, "hr", false, "view size in humar readable format")
		flag.BoolVar(&hrB, "humanfriendly", false, "view size in humar readable format")

		// Tree recursive format
		flag.BoolVar(&tB, "t", false, "view tree structure ( recursive lookup )")
		flag.BoolVar(&tB, "tree", false, "view tree structure ( recursive lookup )")
	*/

	// Set help
	/**
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Example viswa Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
	*/

	//Parse the flag
	flag.Parse()

	decor := longB || drB || fB || hrB || tB
	flags := flagdef{longB, drB, fB, hrB, tB}
	dir := "."
	if flag.NArg() > 0 {
		dir = flag.Args()[0]
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
}
