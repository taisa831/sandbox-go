package ioutil

import "testing"

func TestReadAll(t *testing.T) {
	str := ReadAll()
	if str != "0123456789" {
		t.Errorf("TestReadAll Error. %s", str)
	}
}

func TestReadFile(t *testing.T) {
	str := ReadFile()
	if str != "0123456789" {
		t.Errorf("TestReadAll Error. %s", str)
	}
}

func TestWriteFile(t *testing.T) {
	str := WriteFile()
	if str != "0123456789" {
		t.Errorf("TestReadAll Error. %s", str)
	}
}

func TestReadDir(t *testing.T) {
	fileInfoList := ReadDir()
	for _, fileInfo := range fileInfoList {
		if !(fileInfo.Name() == "src.txt" || fileInfo.Name() == "dst.txt") {
			t.Errorf("TestReadDir Error. %s", fileInfo.Name())
		}
	}
}

func TestNopCloserClose(t *testing.T) {
	if NopCloserClose() != nil {
		t.Errorf("TestNopCloserClose Error. %v", NopCloserClose())
	}
}

func TestDevNullWrite(t *testing.T) {
	if DevNullWrite() != 1024 {
		t.Errorf("TestDevNullWrite Error. %v", DevNullWrite())
	}
}
