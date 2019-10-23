package strings

import "strings"

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