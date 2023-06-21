/*
@File   : deletefilesolderthan.go
@Author : pan
@Time   : 2023-06-21 11:13:11
*/
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type FileFilters struct {
	OlderThan    time.Duration
	Prefix       string
	Suffix       string
	RegexPattern string
	CustomCheck  func(filename string) bool
	Callback     func(filename string) error
}

func DeleteFilesOlderThan(folder string, filter FileFilters) error {
	startScan := time.Now()
	return filepath.WalkDir(folder, func(osPathname string, de fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if osPathname == "" {
			return nil
		}
		if de.IsDir() {
			return nil
		}
		fileInfo, err := os.Stat(osPathname)
		if err != nil {
			return nil
		}
		fileName := fileInfo.Name()
		if filter.Prefix != "" && !strings.HasPrefix(fileName, filter.Prefix) {
			return nil
		}
		if filter.Suffix != "" && !strings.HasSuffix(fileName, filter.Suffix) {
			return nil
		}
		if filter.RegexPattern != "" {
			regex, err := regexp.Compile(filter.RegexPattern)
			if err != nil {
				return err
			}
			if !regex.MatchString(fileName) {
				return nil
			}
		}
		if filter.CustomCheck != nil && !filter.CustomCheck(osPathname) {
			return nil
		}
		if fileInfo.ModTime().Add(filter.OlderThan).Before(startScan) {
			if filter.Callback != nil {
				return filter.Callback(osPathname)
			} else {
				os.RemoveAll(osPathname)
			}
		}
		return nil
	},
	)
}

// 例题怎么赋值结构体里函数的值
// func main() {
// 	filefilter := FileFilters{
// 		Prefix: "test",
// 		Suffix: "txt",
// 		// CustomCheck: func(filename string) bool {
// 		// 	return true
// 		// },
// 		// Callback: func(filename string) error {
// 		// 	return nil
// 		// },
// 	}
// 	result := DeleteFilesOlderThan(`D:\GoProjects\src\intelligentlibrary\ai-go\gotutorial\examples\tests\test.txt`, filefilter)
// 	fmt.Println(result)
// }

func main() {
	fo, err := os.MkdirTemp("", "")
	if err != nil {
		fmt.Printf("mkdirtemp err:%v", err)
	}
	createFile := func() string {
		fi, err := os.CreateTemp(fo, "")
		if err != nil {
			fmt.Printf("create Temp err:%v", err)
		}
		fName := fi.Name()
		fi.Close()
		return fName
	}
	fname := createFile()
	fileInfo, _ := os.Stat(fname)
	filefilter := FileFilters{
		OlderThan: time.Duration(0 * time.Second),
		Prefix:    fileInfo.Name(),
		// CustomCheck: func(filename string) bool {
		// 	return true
		// },
		// Callback: func(filename string) error {
		// 	return nil
		// },
	}
	result := DeleteFilesOlderThan(fo, filefilter)
	fmt.Println(result)
}
