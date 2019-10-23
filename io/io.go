package io

import (
	"io"
	"io/ioutil"
	"os"
)

func WriteString() string {
	file, _ := os.Create("testdata/src.txt")
	_, _ = io.WriteString(file, "0123456789")
	b, _ := ioutil.ReadFile("testdata/src.txt")
	return string(b)
}

func ReadAtLeast() string {
	file, _ := os.Open("testdata/src.txt")
	buf := make([]byte, 10)
	n, _ := io.ReadAtLeast(file, buf, 2)
	return string(buf[:n])
}

func ReadFull() string {
	file, _ := os.Open("testdata/src.txt")
	buf := make([]byte, 4)
	n, _ := io.ReadFull(file, buf)
	return string(buf[:n])
}

func CopyN() string {
	srcFile, _ := os.Open("testdata/src.txt")
	dstFile, _ := os.Create("testdata/dst.txt")
	written, _ := io.CopyN(dstFile, srcFile, 3)

	file, _ := os.Open("testdata/dst.txt")
	buf := make([]byte, written)
	n, _ := file.Read(buf)
	return string(buf[:n])
}

func Copy() string {
	srcFile, _ := os.Open("testdata/src.txt")
	dstFile, _ := os.Create("testdata/dst.txt")
	written, _ := io.Copy(dstFile, srcFile)

	buf := make([]byte, written)
	file, _ := os.Open("testdata/dst.txt")
	n, _ := file.Read(buf)
	return string(buf[:n])
}

func CopyBuffer() string {
	srcFile, _ := os.Open("testdata/src.txt")
	dstFile, _ := os.Create("testdata/dst.txt")
	written, _ := io.CopyBuffer(dstFile, srcFile, make([]byte, 3))

	buf := make([]byte, written)
	file, _ := os.Open("testdata/dst.txt")
	n, _ := file.Read(buf)
	return string(buf[:n])
}

func LimitReaderRead() string {
	srcFile, _ := os.Open("testdata/src.txt")
	limitedReader := io.LimitedReader{srcFile, 5}
	buf := make([]byte, 3)
	n, _ := limitedReader.Read(buf)
	return string(buf[:n])
}

func LimitReaderRead2() string {
	srcFile, _ := os.Open("testdata/src.txt")
	limitedReader := io.LimitedReader{srcFile, 5}
	buf := make([]byte, 8)
	n, _ := limitedReader.Read(buf)
	return string(buf[:n])
}

func SectionReaderRead() string {
	file, _ := os.Open("testdata/src.txt")
	sectionReader := io.NewSectionReader(file, 3, 5)
	buf := make([]byte, 3)
	n, _ := sectionReader.Read(buf)
	return string(buf[:n])
}

func SectionReaderSeek() string {
	file, _ := os.Open("testdata/src.txt")
	sectionReader := io.NewSectionReader(file, 5, 10)
	buf := make([]byte, 3)
	n, _ := sectionReader.Read(buf)
	_, _ = sectionReader.Seek(0, 0)
	return string(buf[:n])
}

func SectionReaderReadAt() string {
	file, _ := os.Open("testdata/src.txt")
	sectionReader := io.NewSectionReader(file, 5, 10)
	buf := make([]byte, 3)
	n, _ := sectionReader.ReadAt(buf, 2)
	return string(buf[:n])
}

func SectionReaderSize() int64 {
	file, _ := os.Open("testdata/src.txt")
	sectionReader := io.NewSectionReader(file, 5, 9)
	return sectionReader.Size()
}

func TeeReaderRead() string {
	srcFile, _ := os.Open("testdata/src.txt")
	dstFile, _ := os.Create("testdata/dst.txt")
	teeReader := io.TeeReader(srcFile, dstFile)
	buf := make([]byte, 3)
	_, _ = teeReader.Read(buf)
	dstFile2, _ := os.Open("testdata/dst.txt")
	n, _ := dstFile2.Read(buf)
	return string(buf[:n])
}