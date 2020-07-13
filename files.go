package main

import (
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
		}
	}
	return contents, paths
}
