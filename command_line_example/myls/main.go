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
// -a show all
// -v version
// -h help
const version = "1.0"

func main() {
	// get current dir
	// dir, err := os.Getwd()

	// parse args option
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version.")
	var showDetail bool
	flag.BoolVar(&showDetail, "l", false, "show list with more detail.")
	var showAll bool
	flag.BoolVar(&showAll, "a", false, "show all files.")
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
		longFormat(fis, showAll)
		return
	}
	ls(fis, showAll)
	return
}

func ls(fis []os.FileInfo, showAll bool) {
	for _, fi := range fis {
		filename := fi.Name()
		if !showAll && strings.HasPrefix(filename, ".") {
			// ignore dot files
			continue
		}
		fmt.Print(filename + " ")
	}
	fmt.Println("")
}

func longFormat(fis []os.FileInfo, showAll bool) {
	var stringFileInfos []*stringFileInfo

	for _, fi := range fis {
		info := stringFileInfo{}

		filename := fi.Name()
		if !showAll && strings.HasPrefix(filename, ".") {
			// ignore dot files
			continue
		}

		// file name
		info.name = filename
		// permission
		info.permission = fmt.Sprintf("%v", fi.Mode())
		// file size
		info.size = fmt.Sprintf("%v", fi.Size())
		// user name
		var s syscall.Stat_t
		syscall.Stat(fi.Name(), &s)
		u, _ := user.LookupId(fmt.Sprintf("%v", s.Uid))
		info.user = u.Username
		// group name
		g, _ := user.LookupGroupId(fmt.Sprintf("%v", s.Gid))
		info.group = g.Name
		// mod time
		modTime := fi.ModTime()
		info.modDateTime = modTime.Format("1 2 15:04")

		stringFileInfos = append(stringFileInfos, &info)
	}

	fillSpaces(stringFileInfos)
	for _, info := range stringFileInfos {
		fmt.Printf("%s %s %s %s %s %s\n", info.permission, info.user, info.group, info.size, info.modDateTime, info.name)
	}
}

// fill spaces for format
func fillSpaces(infos []*stringFileInfo) {
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
	var spaceSize int
	for _, info := range infos {
		// name
		spaceSize = maxLengthName - len(info.name)
		info.name += strings.Repeat(" ", spaceSize)
		// size
		spaceSize = maxLengthSize - len(info.size)
		info.size += strings.Repeat(" ", spaceSize)
		// user name
		spaceSize = maxLengthUser - len(info.user)
		info.user += strings.Repeat(" ", spaceSize)
		// group name
		spaceSize = maxLengthGroup - len(info.group)
		info.group += strings.Repeat(" ", spaceSize)
		// mod date
		spaceSize = maxLengthDate - len(info.modDateTime)
		info.modDateTime += strings.Repeat(" ", spaceSize)
	}
}

type stringFileInfo struct {
	name        string
	permission  string
	size        string
	user        string
	group       string
	modDateTime string
}
