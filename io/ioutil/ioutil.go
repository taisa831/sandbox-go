package ioutil

import (
	"io/ioutil"
	"os"
)

func ReadAll() string {
	file, _ := os.Open("testdata/src.txt")
	b, _ := ioutil.ReadAll(file)
	return string(b)
}

func ReadFile() string {
	b, _ := ioutil.ReadFile("testdata/src.txt")
	return string(b)
}

func WriteFile() string {
	b := []byte("0123456789")
	_ = ioutil.WriteFile("testdata/dst.txt", b, os.ModePerm)
	b, _ = ioutil.ReadFile("testdata/dst.txt")
	return string(b)
}

func ReadDir() []os.FileInfo {
	fileInfo, _ := ioutil.ReadDir("testdata")
	return fileInfo
}