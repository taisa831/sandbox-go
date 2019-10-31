package io

import (
	"io"
	"io/ioutil"
	"os"
)

// writerに文字列を書き込む
func WriteString() string {
	file, _ := os.Create("testdata/src.txt")
	_, _ = io.WriteString(file, "0123456789")
	b, _ := ioutil.ReadFile("testdata/src.txt")
	return string(b)
}

// 指定した値は最低限読み込みバッファー分すべて読み込む
func ReadAtLeast(b int, min int) (string, error) {
	file, _ := os.Open("testdata/src.txt")
	buf := make([]byte, b)
	n, err := io.ReadAtLeast(file, buf, min)
	return string(buf[:n]), err
}

// 指定したバッファー分すべて読み込む
func ReadFull() string {
	file, _ := os.Open("testdata/src.txt")
	buf := make([]byte, 5)
	n, _ := io.ReadFull(file, buf)
	return string(buf[:n])
}

// writerへ指定した値をコピーする
func CopyN() string {
	srcFile, _ := os.Open("testdata/src.txt")
	dstFile, _ := os.Create("testdata/dst.txt")
	written, _ := io.CopyN(dstFile, srcFile, 5)
	file, _ := os.Open("testdata/dst.txt")
	buf := make([]byte, written)
	n, _ := file.Read(buf)
	return string(buf[:n])
}

// readerをwriterにコピーする
func Copy() string {
	srcFile, _ := os.Open("testdata/src.txt")
	dstFile, _ := os.Create("testdata/dst.txt")
	written, _ := io.Copy(dstFile, srcFile)
	buf := make([]byte, written)
	file, _ := os.Open("testdata/dst.txt")
	n, _ := file.Read(buf)
	return string(buf[:n])
}

// readerをwriterに指定したバッファー分コピーする
func CopyBuffer() string {
	srcFile, _ := os.Open("testdata/src.txt")
	dstFile, _ := os.Create("testdata/dst.txt")
	written, _ := io.CopyBuffer(dstFile, srcFile, make([]byte, 5))
	buf := make([]byte, written)
	file, _ := os.Open("testdata/dst.txt")
	n, _ := file.Read(buf)
	return string(buf[:n])
}

// 値を制限したreaderを取得する
func LimitReaderRead() string {
	srcFile, _ := os.Open("testdata/src.txt")
	limitedReader := io.LimitedReader{srcFile, 5}
	buf := make([]byte, 3)
	n, _ := limitedReader.Read(buf)
	return string(buf[:n])
}

// 値を制限したreaderを取得し読み込む
func LimitReaderRead2() string {
	srcFile, _ := os.Open("testdata/src.txt")
	limitedReader := io.LimitedReader{srcFile, 5}
	buf := make([]byte, 8)
	n, _ := limitedReader.Read(buf)
	return string(buf[:n])
}

// 指定した値でセクション分けしたreaderを取得し読み込む
func SectionReaderRead() string {
	file, _ := os.Open("testdata/src.txt")
	sectionReader := io.NewSectionReader(file, 3, 5)
	buf := make([]byte, 3)
	n, _ := sectionReader.Read(buf)
	return string(buf[:n])
}

// セクションリーダーの読み込み位置を変更する
func SectionReaderSeek() string {
	file, _ := os.Open("testdata/src.txt")
	sectionReader := io.NewSectionReader(file, 5, 10)
	buf := make([]byte, 3)
	n, _ := sectionReader.Read(buf)
	_, _ = sectionReader.Seek(0, 0)
	return string(buf[:n])
}

// セクションリーダーを指定した位置から読み込む
func SectionReaderReadAt() string {
	file, _ := os.Open("testdata/src.txt")
	sectionReader := io.NewSectionReader(file, 5, 10)
	buf := make([]byte, 3)
	n, _ := sectionReader.ReadAt(buf, 2)
	return string(buf[:n])
}

// セクションリーダーのサイズを取得する
func SectionReaderSize() int64 {
	file, _ := os.Open("testdata/src.txt")
	sectionReader := io.NewSectionReader(file, 5, 9)
	return sectionReader.Size()
}

// readerとwriterを渡しておいて読み込んだら同時にwriterにも書き込む
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