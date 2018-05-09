package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"strings"
	"syscall"
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
	var names []string
	for _, fi := range fis {
		names = append(names, fi.Name())
	}
	fillSpaces(names)

	for _, fi := range fis {
		// TODO show file or directory
		// if file, -
		// if dir, d

		// permission
		fmt.Printf("%v ", fi.Mode())
		fmt.Print(" ")

		// user name, group name
		// fmt.Print(s.Uid)
		// fmt.Print(" ")
		// fmt.Print(s.Gid)
		// fmt.Print(" ")
		var s syscall.Stat_t
		syscall.Stat(fi.Name(), &s)
		u, _ := user.LookupId(fmt.Sprintf("%v", s.Uid))
		fmt.Print(u.Username)
		fmt.Print(" ")
		g, _ := user.LookupGroupId(fmt.Sprintf("%v", s.Gid))
		fmt.Print(g.Name)
		fmt.Print(" ")
		fmt.Printf("%v ", fi.Size())
		fmt.Print(" ")

		// mod time
		modTime := fi.ModTime()
		// fmt.Print(modTime.Month() + " " modTime.Day() + " " + modTime.Day())
		// m d hh24:mi
		fmt.Print(modTime.Format("1 2 15:04"))
		fmt.Print(" ")

		// file name
		fmt.Print(fi.Name() + " ")
		fmt.Println("")

		info := StringFileInfo{}
		info.name = fi.Name()
		info.permission = fmt.Sprintf("%v", fi.Mode())
		info.size = fmt.Sprintf("%v", fi.Size())
		info.user = u.Username
		info.group = g.Name
		info.modDateTime = modTime.Format("1 2 15:04")
	}
}

// fill spaces for format
func fillSpaces(texts []string) {
	var maxLength int
	for _, text := range texts {
		if maxLength < len(text) {
			maxLength = len(text)
		}
	}

	for _, text := range texts {
		spaceSize := maxLength - len(text)
		spaces := strings.Repeat(" ", spaceSize+1)
		text = text + spaces
		fmt.Println(":" + text + ":")
	}
}

type StringFileInfo struct {
	name        string
	permission  string
	size        string
	user        string
	group       string
	modDateTime string
}
