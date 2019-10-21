package io

import (
	"testing"
)

func TestWriteString(t *testing.T) {
	str := WriteString()
	if str != "0123456789" {
		t.Errorf("WriteString Error %s", str)
	}
}

func TestReadAtLeast(t *testing.T) {
	str := ReadAtLeast()
	if str != "012" {
		t.Errorf("TestReadAtLeast Error %s", str)
	}
}

func TestReadFull(t *testing.T) {
	str := ReadFull()
	if str != "0123" {
		t.Errorf("TestReadFull Error %s", str)
	}
}

func TestCopyN(t *testing.T) {
	str := CopyN()
	if str != "012" {
		t.Errorf("TestCopyN Error %s", str)
	}
}

func TestCopy(t *testing.T) {
	n, _ := Copy()
	if n != 10 {
		t.Errorf("TestCopy Error %d", n)
	}
}

func TestCopyBuffer(t *testing.T) {
	n, _ := CopyBuffer()
	if n != 10 {
		t.Errorf("TestCopyBuffer Error %d", n)
	}
}

func TestLimitReaderRead(t *testing.T) {
	n, _ := LimitReaderRead()
	if n != 2 {
		t.Errorf("TestLimitReaderRead Error %d" , n)
	}
}

func TestSectionReaderRead(t *testing.T) {
	n, _ := SectionReaderRead()
	if n != 3 {
		t.Errorf("TestSectionReaderRead Error %d", n)
	}
}

func TestSectionReaderSeek(t *testing.T) {
	n, _ := SectionReaderSeek()
	if n != 2 {
		t.Errorf("TestSectionReaderSeek Error %d", n)
	}
}

func TestSectionReaderReadAt(t *testing.T) {
	n, _ := SectionReaderReadAt()
	if n != 3 {
		t.Errorf("TestSectionReaderReadAt Error %d", n)
	}
}

func TestSectionReaderSize(t *testing.T) {
	size := SectionReaderSize()
	if size != 5 {
		t.Errorf("TestSectionReaderSize Error %d", size)
	}
}

func TestTeeReaderRead(t *testing.T) {
	n, _ := TeeReaderRead()
	if n != 3 {
		t.Errorf("TestTeeReaderRead Error %d", n)
	}
}