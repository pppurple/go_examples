package main

import (
	"flag"
	"fmt"
)

func main() {
	var showList bool
	flag.BoolVar(&showList, "list", false, "show list.")
	var fileName string
	flag.StringVar(&fileName, "f", "dummy.txt", "file name.")
	var maxSize int
	flag.IntVar(&maxSize, "s", 0, "max size.")
	flag.Parse()

	fmt.Printf("bool: %t\n", showList)
	fmt.Printf("string: %s\n", fileName)
	fmt.Printf("int: %d\n", maxSize)
}
