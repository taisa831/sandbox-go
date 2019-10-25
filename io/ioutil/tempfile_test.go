package ioutil

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestTempFile(t *testing.T) {
	f := TempFile()
	if f.Name() == "" {
		t.Errorf("TestTempFile error %s", f.Name())
	}
	defer os.RemoveAll("testdata/tempfile")
}

func TestTempDir(t *testing.T) {
	name := TempDir()
	if name == "" {
		t.Errorf("TestTempDir error %s", name)
	}
	fileInfoList, _ := ioutil.ReadDir("testdata")
	for _, fileInfo := range fileInfoList {
		if fileInfo.IsDir() {
			os.RemoveAll("testdata/" + fileInfo.Name())
		}
	}
}