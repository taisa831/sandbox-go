package ioutil

import (
	"io/ioutil"
	"os"
)

func TempFile() *os.File {
	f , _ := ioutil.TempFile("testdata/tempfile", "test")
	return f
}

func TempDir() string {
	name, _ := ioutil.TempDir("testdata", "test")
	return name
}