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

	var stringFileInfos []StringFileInfo

	for _, fi := range fis {
		// TODO show file or directory
		// if file, -
		// if dir, d

		// permission
		fmt.Printf("%v ", fi.Mode())

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
		stringFileInfos = append(stringFileInfos, info)
	}

	fill(stringFileInfos)
	for _, info := range stringFileInfos {
		fmt.Printf("%s %s %s %s %s %s\n", info.permission, info.user, info.group, info.size, info.modDateTime, info.name)
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

func getMaxLength(texts []string) {

}

func fill(infos []StringFileInfo) {
	var maxLengthName int
	var maxLengthPerm int
	var maxLengthSize int
	var maxLengthUser int
	var maxLengthGroup int
	var maxLengthDate int
	for _, info := range infos {
		if maxLengthName < len(info.name) {
			maxLengthName = len(info.name)
		}
		if maxLengthPerm < len(info.permission) {
			maxLengthPerm = len(info.permission)
		}
		if maxLengthSize < len(info.size) {
			maxLengthSize = len(info.size)
		}
		if maxLengthUser < len(info.user) {
			maxLengthUser = len(info.user)
		}
		if maxLengthGroup < len(info.group) {
			maxLengthGroup = len(info.group)
		}
		if maxLengthDate < len(info.modDateTime) {
			maxLengthDate = len(info.modDateTime)
		}
	}
	fmt.Println(maxLengthName)
	fmt.Println(maxLengthPerm)
	fmt.Println(maxLengthSize)
	fmt.Println(maxLengthUser)
	fmt.Println(maxLengthGroup)
	fmt.Println(maxLengthDate)
	var spaceSize int
	for _, info := range infos {
		// name
		spaceSize = maxLengthName - len(info.name)
		info.name = info.name + strings.Repeat(" ", spaceSize+1)
		// permission
		// spaceSize = maxLengthPerm - len(info.permission)
		// info.permission = info.permission + strings.Repeat(" ", spaceSize+1)
		// size
		spaceSize = maxLengthSize - len(info.size)
		info.size = info.size + strings.Repeat(" ", spaceSize+1)
		fmt.Println(":" + info.size + ":")
		// user name
		spaceSize = maxLengthUser - len(info.user)
		info.user = info.user + strings.Repeat(" ", spaceSize+1)
		// group name
		spaceSize = maxLengthGroup - len(info.group)
		info.group = info.group + strings.Repeat(" ", spaceSize+1)
		// mod date
		spaceSize = maxLengthDate - len(info.modDateTime)
		info.modDateTime = info.modDateTime + strings.Repeat(" ", spaceSize+1)
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
