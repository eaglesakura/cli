package utils

import (
	"io/ioutil"
	"os"
	"strconv"
)

func Atoi(ascii string) int {
	result, err := strconv.Atoi(ascii)
	if err != nil {
		return 0
	} else {
		return result
	}
}

/*
	list directories in 'path'.
*/
func ListDirectories(path string) []os.FileInfo {
	files, _ := ioutil.ReadDir(path)
	var result []os.FileInfo

	for _, info := range files {
		if info.IsDir() {
			result = append(result, info)
		}
	}
	return result
}

/*
	list files in 'path'.
*/
func ListFiles(path string) []os.FileInfo {
	files, _ := ioutil.ReadDir(path)
	var result []os.FileInfo
	for _, info := range files {
		if !info.IsDir() {
			result = append(result, info)
		}
	}
	return result
}
