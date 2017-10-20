package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("My first golang code...")
	fmt.Println("Implementing ls in golang...")
	f, _ := os.OpenFile(".", os.O_RDONLY, 0666)
	files, _ := f.Readdirnames(0)
	for _, file := range files {
		fmt.Println(file)
	}
}
