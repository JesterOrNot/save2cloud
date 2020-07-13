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
		PrintError("Path does not exist.", 127)
	}
	for _, file := range files {
		fileContents, err := ioutil.ReadFile(path + "/" + file.Name())
		if err != nil {
			PrintError("Could not read file.", 1)
		}
		contents = append(contents, string(fileContents))
		paths = append(paths, file.Name())
	}
	return contents, paths
}
