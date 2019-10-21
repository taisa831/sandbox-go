package io

import (
	"testing"
)

func TestWriteString(t *testing.T) {
	n, _ := WriteString()
	if n != 10 {
		t.Error("WriteString Error")
	}
}

func TestReadAtLeast(t *testing.T) {
	n, _ := ReadAtLeast()
	if n != 3 {
		t.Errorf("TestReadAtLeast Error %d", n)
	}
}

func TestReadFull(t *testing.T) {
	n, _ := ReadFull()
	if n != 4 {
		t.Errorf("TestReadFull Error %d", n)
	}
}

func TestCopyN(t *testing.T) {
	n, _ := CopyN()
	if n != 3 {
		t.Errorf("TestCopyN Error %d", n)
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