package io

import (
	"io"
	"os"
)

func WriteString() (int, error) {
	file, _ := os.Create("testdata/src.txt")
	return io.WriteString(file, "0123456789")
}

func ReadAtLeast() (int, error) {
	file, _ := os.Open("testdata/src.txt")
	buf := make([]byte, 3)
	return io.ReadAtLeast(file, buf, 3)
}

func ReadFull() (int, error) {
	file, _ := os.Open("testdata/src.txt")
	buf := make([]byte, 4)
	return io.ReadFull(file, buf)
}

func CopyN() (int64, error) {
	srcFile, _ := os.Open("testdata/src.txt")
	dstFile, _ := os.Create("testdata/dst.txt")
	return io.CopyN(dstFile, srcFile, 3)
}

func Copy() (int64, error) {
	srcFile, _ := os.Open("testdata/src.txt")
	dstFile, _ := os.Create("testdata/dst.txt")
	return io.Copy(dstFile, srcFile)
}

func CopyBuffer() (int64, error) {
	srcFile, _ := os.Open("testdata/src.txt")
	dstFile, _ := os.Create("testdata/dst.txt")
	return io.CopyBuffer(dstFile, srcFile, make([]byte, 3))
}

func LimitReaderRead() (int, error) {
	file, _ := os.Open("testdata/src.txt")
	limitedReader := io.LimitedReader{file, 2}
	buf := make([]byte, 3)
	return limitedReader.Read(buf)
}

func SectionReaderRead() (int, error) {
	file, _ := os.Open("testdata/src.txt")
	sectionReader := io.NewSectionReader(file, 3, 5)
	buf := make([]byte, 3)
	n, _ := sectionReader.Read(buf)
	return n, nil
}

func SectionReaderSeek() (int64, error) {
	file, _ := os.Open("testdata/src.txt")
	sectionReader := io.NewSectionReader(file, 3, 5)
	return sectionReader.Seek(2, 1)
}

func SectionReaderReadAt() (int, error) {
	file, _ := os.Open("testdata/src.txt")
	sectionReader := io.NewSectionReader(file, 3, 5)
	buf := make([]byte, 3)
	return sectionReader.ReadAt(buf, 2)
}

func SectionReaderSize() int64 {
	file, _ := os.Open("testdata/src.txt")
	sectionReader := io.NewSectionReader(file, 3, 5)
	return sectionReader.Size()
}

func TeeReaderRead() (int, error) {
	r, _ := os.Open("testdata/src.txt")
	w, _ := os.Create("testdata/dst.txt")
	teeReader := io.TeeReader(r, w)
	buf := make([]byte, 3)
	return teeReader.Read(buf)
}