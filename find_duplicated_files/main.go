package main

import (
	"flag"
	"log"
	"os"
	"regexp"
)

const (
	codeOK    = 0
	codeError = 1
)

func main() {
	execute()
}

func execute() int {
	flag.Parse()
	dirName := flag.Arg(0)
	println("dirName:" + dirName)
	if dirName == "" {
		dirName = "."
	}

	f, err := os.Open(dirName)
	if err != nil {
		log.Fatal("Cannot open dir. Please input correct dir name:")
	}
	defer f.Close()

	fis, err := f.Readdir(0)
	if err != nil {
		log.Fatal("Cannot read dir. Please input correct dir name:")
	}

	ignorePattern := [2]string{`(1)$`, `_1\..+$`}

	// fileMap := map[string]int{}
	var files = []string{}
	for _, fi := range fis {
		filename := fi.Name()
		println("filename:" + filename)

		// ignore suffix patterns
		for _, pattern := range ignorePattern {
			r := regexp.MustCompile(pattern)
			if r.MatchString(filename) {
				println("match!!:" + filename)
				files = append(files, filename)
				// fileMap[filename]++
			}
		}

		/*
			_, exist := fileMap[filename]
			if exist {
				fileMap[filename]++
			} else {
				fileMap[filename] = 0
			}
		*/
	}

	for _, file := range files {
		println(file)
	}

	return codeOK
}
