package main

import (
	"fmt"
	"io/ioutil"
)

// GetFilesInDir gets the contents of all files in a directory
func GetFilesInDir(path string) ([]string, []string) {
	var contents []string
	var paths []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		PrintError(err.Error(), 127)
	}
	for _, file := range files {
		if !file.IsDir() {
			fileContents, err := ioutil.ReadFile(path + "/" + file.Name())
			if err != nil {
				PrintError(err.Error(), 1)
			}
			contents = append(contents, string(fileContents))
			paths = append(paths, file.Name())
		} else {
			dirContents, dirPaths := GetFilesInDir(path + "/" + file.Name())
			for i, dirPath := range dirPaths {
				dirPaths[i] = file.Name() + "/" + dirPath
			}
			contents = append(append(contents, ""), dirContents...)
			paths = append(append(paths, fmt.Sprintf("%s/", file.Name())), dirPaths...)
		}
	}
	return contents, paths
}
