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
const (
	version   = "1.0"
	codeOk    = 0
	codeError = 1
)

func main() {
	os.Exit(executeCommand())
}

func executeCommand() int {
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
		return codeOk
	}

	// dir name
	dirName := flag.Arg(0)
	if dirName == "" {
		dirName = "."
	}

	/*
		files, err := ioutil.ReadDir("./")
		if err != nil {
			log.Fatal(err)
		}
	*/

	f, err := os.Open(dirName)
	if err != nil {
		log.Fatal("Cannot open dir. Please input correct dir name:")
	}
	defer f.Close()

	fis, err := f.Readdir(0)
	if err != nil {
		log.Fatal("Cannot open dir. Please input correct dir name:")
	}

	if showDetail {
		longFormat(fis, showAll)
		return codeOk
	}
	ls(fis, showAll)
	return codeOk
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
	var maxLengthMonth int
	var maxLengthDay int
	var maxLengthTime int
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
		// split "M D hh:mi"
		dateTime := strings.Split(info.modDateTime, " ")
		if maxLengthMonth < len(dateTime[0]) {
			maxLengthMonth = len(dateTime[0])
		}
		if maxLengthDay < len(dateTime[1]) {
			maxLengthDay = len(dateTime[1])
		}
		if maxLengthTime < len(dateTime[2]) {
			maxLengthTime = len(dateTime[2])
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
		// mod date time (month + day + time)
		dateTime := strings.Split(info.modDateTime, " ")
		spaceSize = maxLengthMonth - len(dateTime[0]) // month
		info.modDateTime = dateTime[0] + strings.Repeat(" ", spaceSize+1)
		spaceSize = maxLengthDay - len(dateTime[1]) // day
		info.modDateTime += dateTime[1] + strings.Repeat(" ", spaceSize+1)
		spaceSize = maxLengthTime - len(dateTime[2]) // time
		info.modDateTime += dateTime[2] + strings.Repeat(" ", spaceSize)
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
