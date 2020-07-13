package main

import (
	"io/ioutil"
	"os"
	"sync"
	"testing"
)

func CreateFileWithContent(path string, content string, wg *sync.WaitGroup) {
	bytes := []byte(content)
	os.Create(path)
	err := ioutil.WriteFile(path, bytes, 0666)
	if err != nil {
		PrintError(err.Error(), 1)
	}
	wg.Done()
}

func TestGetFilesInDir(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	os.Mkdir("test", 0777)
	go CreateFileWithContent("./test/a.txt", "Hello World From a.txt", &wg)
	go CreateFileWithContent("./test/b.txt", "Hello World From b.txt", &wg)
	go CreateFileWithContent("./test/c.txt", "Hello World From c.txt", &wg)
	wg.Wait()
	contents, _ := GetFilesInDir("test")
	if !Contains(contents, "Hello World From a.txt") {
		PrintError("Failed", 1)
	} else if !Contains(contents, "Hello World From b.txt") {
		PrintError("Failed", 1)
	} else if !Contains(contents, "Hello World From c.txt") {
		PrintError("Failed", 1)
	}
}
