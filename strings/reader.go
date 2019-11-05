package strings

import (
	"bytes"
	"strings"
)

func Len() int {
	reader := strings.NewReader("0123456789")
	return reader.Len()
}

func Size() int64 {
	reader := strings.NewReader("0123456789")
	return reader.Size()
}

func Read() string {
	reader := strings.NewReader("0123456789")
	buf := make([]byte, 3)
	n, _ := reader.Read(buf)
	return string(buf[:n])
}

func ReadAt() string {
	reader := strings.NewReader("0123456789")
	buf := make([]byte, 3)
	n, _ := reader.ReadAt(buf, 2)
	return string(buf[:n])
}

func ReadByte() byte {
	reader := strings.NewReader("0123456789")
	b, _ := reader.ReadByte()
	return b
}

func UnreadByte() string {
	reader := strings.NewReader("0123456789")
	_ = reader.UnreadByte()
	b, _ := reader.ReadByte()
	return string(b)
}

// UTF-8エンコードされたUnicode文字を一文字読み込みruneにセットします。読み込んだバイト数はsizeにセットされます。
func ReadRune() (rune, int) {
	reader := strings.NewReader("0123456789")
	ch, size, _ := reader.ReadRune()
	return ch, size
}

func UnreadRune() (rune, int) {
	reader := strings.NewReader("0123456789")
	_ = reader.UnreadRune()
	ch, size, _ := reader.ReadRune()
	return ch, size
}

func Seek() string {
	reader := strings.NewReader("0123456789")
	b := make([]byte, 10)
	_, _ = reader.Read(b)
	n, _ := reader.Read(b)
	return string(b[:n])
}

func Seek2() string {
	reader := strings.NewReader("0123456789")
	b := make([]byte, 10)
	_, _ = reader.Read(b)
	reader.Reset("012")
	n, _ := reader.Read(b)
	return string(b[:n])
}

func WriteTo() int64 {
	reader := strings.NewReader("0123456789")
	buf := new(bytes.Buffer)
	n, _ := reader.WriteTo(buf)
	return n
}

func Reset() string {
	reader := strings.NewReader("0123456789")
	b := make([]byte, 10)
	_, _ = reader.Read(b)
	n, _ := reader.Read(b)
	return string(b[:n])
}

func Reset2() string {
	reader := strings.NewReader("0123456789")
	b := make([]byte, 10)
	_, _ = reader.Read(b)
	reader.Reset("012")
	n, _ := reader.Read(b)
	return string(b[:n])
}