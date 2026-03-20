package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/AJRDRGZ/fileinfo"
	"github.com/fatih/color"
	"golang.org/x/exp/constraints"
)

const Windows = "windows"
const Linux = "linux"
const Mac = "mac"

func main() {
	// filter patterns
	flagPattern := flag.String("p", "", "filter by pattern")
	flagAll := flag.Bool("a", false, "all files including hidden files")
	flagNumberRecords := flag.Int("n", 0, "number of records")

	// orden flags
	hasOrderByTime := flag.Bool("t", false, "sort by time, oldest first")
	hasOrderBySize := flag.Bool("s", false, "sort by size, smallest first")
	hasOrderReverse := flag.Bool("r", false, "reverse order while sorting")

	flag.Parse()
	dirPath := flag.Arg(0)
	if dirPath == "" {
		dirPath = "."
	}

	dirs, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading directory: %v\n", err)
		os.Exit(1)
	}

	var pattern *regexp.Regexp
	if *flagPattern != "" {
		pattern, err = regexp.Compile("(?i)" + *flagPattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid pattern: %v\n", err)
			os.Exit(1)
		}
	}

	fs := []file{}

	for _, dir := range dirs {
		isHidden := isHidden(dir.Name(), dirPath)

		if isHidden && !*flagAll {
			continue
		}

		if pattern != nil && !pattern.MatchString(dir.Name()) {
			continue
		}

		f, err := getFile(dir, isHidden)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error getting file info: %v\n", err)
			os.Exit(1)
		}

		fs = append(fs, f)
	}

	if !*hasOrderBySize || !*hasOrderByTime {
		orderByName(fs, *hasOrderReverse)
	}

	if *hasOrderBySize && !*hasOrderByTime {
		orderBySize(fs, *hasOrderReverse)
	}

	if *hasOrderByTime {
		orderByTime(fs, *hasOrderReverse)
	}

	if *flagNumberRecords == 0 || *flagNumberRecords > len(fs) {
		*flagNumberRecords = len(fs)
	}

	printList(fs, *flagNumberRecords)
}

func mySort[T constraints.Ordered](i, j T, isReverse bool) bool {
	if isReverse {
		return i > j
	}
	return i < j
}

func orderByTime(files []file, isReverse bool) {
	sort.SliceStable(files, func(i, j int) bool {
		return mySort(
			files[i].modificationTime.Unix(),
			files[j].modificationTime.Unix(),
			isReverse,
		)
	})
}

func orderByName(files []file, isReverse bool) {
	sort.SliceStable(files, func(i, j int) bool {
		return mySort(
			strings.ToLower(files[i].name),
			strings.ToLower(files[j].name),
			isReverse,
		)
	})
}

func orderBySize(files []file, isReverse bool) {
	sort.SliceStable(files, func(i, j int) bool {
		return mySort(
			files[i].size,
			files[j].size,
			isReverse,
		)
	})
}

func printList(fs []file, nRecords int) {
	for _, file := range fs[:nRecords] {
		style := mapStyleByFileType[file.fileType]
		fmt.Printf("%s %s %s %10d %s %s %s %s %s\n", file.mode, file.userName,
			file.groupName, file.size, file.modificationTime.Format(time.DateTime),
			style.icon, setColor(file.name, style.color), style.symbol, markHidden(file.isHidden))
	}
}

func getFile(dir fs.DirEntry, isHidden bool) (file, error) {
	info, err := dir.Info()
	if err != nil {
		return file{}, fmt.Errorf("dir.Info(): %w", err)
	}

	userName, groupName := fileinfo.GetUserAndGroup(info.Sys())

	f := file{
		name:             dir.Name(),
		fileType:         fileRegular,
		isDir:            dir.IsDir(),
		isHidden:         isHidden,
		userName:         userName,
		groupName:        groupName,
		size:             info.Size(),
		modificationTime: info.ModTime(),
		mode:             info.Mode().String(),
	}

	setFile(&f)

	return f, nil
}

func setFile(f *file) {
	switch {
	case isLink(*f):
		f.fileType = fileLink
	case f.isDir:
		f.fileType = fileDirectory
	case isExecutable(*f):
		f.fileType = fileExecutable
	case isCompress(*f):
		f.fileType = fileCompress
	case isImage(*f):
		f.fileType = fileImage
	default:
		f.fileType = fileRegular
	}
}

func setColor(nameFile string, styleColor color.Attribute) string {
	switch styleColor {
	case color.FgBlue:
		return blue(nameFile)
	case color.FgGreen:
		return green(nameFile)
	case color.FgRed:
		return red(nameFile)
	case color.FgMagenta:
		return magenta(nameFile)
	case color.FgCyan:
		return cyan(nameFile)
	}
	return nameFile
}

func isLink(f file) bool {
	return strings.HasPrefix(strings.ToUpper(f.mode), "L")
}

func isExecutable(f file) bool {
	if runtime.GOOS == Windows {
		return strings.HasSuffix(f.mode, exe)
	}
	return strings.Contains(f.mode, "x")
}

func isCompress(f file) bool {
	return strings.HasSuffix(f.name, zip) ||
		strings.HasSuffix(f.name, gz) ||
		strings.HasSuffix(f.name, tar) ||
		strings.HasSuffix(f.name, rar) ||
		strings.HasSuffix(f.name, deb)
}

func isImage(f file) bool {
	return strings.HasSuffix(f.name, png) ||
		strings.HasSuffix(f.name, jpg) ||
		strings.HasSuffix(f.name, gif)
}

func isHidden(fileName, basePath string) bool {
	filePath := fileName
	if runtime.GOOS == Windows {
		filePath = path.Join(basePath, fileName)
	}
	return fileinfo.IsHidden(filePath)
}

func markHidden(isHidden bool) string {
	if !isHidden {
		return ""
	}
	return yellow("∅")
}

// go mod init main
