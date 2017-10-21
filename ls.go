package main

import (
	"fmt"
	"os"
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
	if len(os.Args) > 1 {
		serveDir(os.Args[1])
	} else {
		serveDir(".")
	}
}
