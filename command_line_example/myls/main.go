package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// -l format long
// -a all
// -t sort time
// -r reverse
// -h
const version = "1.0"

func main() {
	// get current dir
	// dir, err := os.Getwd()

	// parse args option
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version.")
	var showDetail bool
	flag.BoolVar(&showDetail, "l", false, "show list with more detail.")
	flag.Parse()

	if showVersion {
		fmt.Println("version: " + version)
		return
	}

	f, err := os.Open(".")
	if err != nil {
		log.Fatal("Please input dir name:")
	}
	defer f.Close()

	fis, err := f.Readdir(0)
	if err != nil {
		log.Fatal("Please input dir name:")
	}

	if showDetail {
		longFormat(fis)
		return
	} else {
		ls(fis)
		return
	}

	for _, fi := range fis {
		fmt.Println(fi.Name())
		fmt.Println(fi.Size())
		fmt.Println(fi.Mode())
		fmt.Println(fi.ModTime())
		fmt.Println(fi.IsDir())
		fmt.Println(fi.Sys())

		if fi.IsDir() {
			// fmt.Println(fi.Name())
		}
	}
}

func ls(fis []os.FileInfo) {
	for _, fi := range fis {
		fmt.Print(fi.Name() + " ")
	}
	fmt.Println("")
}

func longFormat(fis []os.FileInfo) {
	for _, fi := range fis {
		fmt.Print(fi.Name() + " ")
		fmt.Println("")
	}
}
