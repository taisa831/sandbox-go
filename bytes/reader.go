package bytes

import "bytes"

func Len() int {
	reader := bytes.NewReader([]byte("0123456789"))
	return reader.Len()
}

func Size() int64 {
	reader := bytes.NewReader([]byte("0123456789"))
	return reader.Size()
}

func Read() string {
	reader := bytes.NewReader([]byte("0123456789"))
	buf := make([]byte, 3)
	n, _ := reader.Read(buf)
	return string(buf[:n])
}