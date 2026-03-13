package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	// filter patterns
	// flagPattern := flag.String("p", "", "filter by pattern")
	// flagAll := flag.Bool("a", false, "all files including hidden files")
	// flagNumerRecords := flag.Int("n", 0, "number of records")

	// orden flags
	// hasOrderByTime := flag.Bool("t", false, "sort by time, oldest first")
	// hasOrderBySize := flag.Bool("s", false, "sort by size, smallest first")
	// hasOrderReverse := flag.Bool("r", false, "reverse order while sorting")

	flag.Parse()
	path := flag.Arg(0)
	if path == "" {
		path = "."
	}

	dirs, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	fs := []file{}

	for _, dir := range dirs {
		f, err := getFile(dir, false)
		if err != nil {
			panic(err)
		}

		fs = append(fs, f)
	}

	fmt.Println(fs)

	// fmt.Println("pattern:", *flagPattern)
	// fmt.Println("all:", *flagAll)
	// fmt.Println("number of records:", *flagNumerRecords)
	// fmt.Println("hasOrderByTime:", *hasOrderByTime)
	// fmt.Println("hasOrderBySize:", *hasOrderBySize)
	// fmt.Println("hasOrderReverse:", *hasOrderReverse)
}

func getFile(dir fs.DirEntry, isHidden bool) (file, error) {
	info, err := dir.Info()
	if err != nil {
		return file{}, fmt.Errorf("dir.Info(): %w", err)
	}

	f := file{
		name:             dir.Name(),
		fileType:         0,
		isDir:            dir.IsDir(),
		isHidden:         isHidden,
		userName:         "123",
		groupName:        "edls",
		size:             info.Size(),
		modificationTime: info.ModTime(),
		mode:             info.Mode().String(),
	}

	return f, nil
}

// go mod init main
