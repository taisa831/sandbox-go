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
	if str != "0123456789" {
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
	str := Copy()
	if str != "0123456789" {
		t.Errorf("TestCopy Error %s", str)
	}
}

func TestCopyBuffer(t *testing.T) {
	str := CopyBuffer()
	if str != "0123456789" {
		t.Errorf("TestCopyBuffer Error %s", str)
	}
}

func TestLimitReaderRead(t *testing.T) {
	str := LimitReaderRead()
	if str != "012" {
		t.Errorf("TestLimitReaderRead Error %s" , str)
	}
}

func TestLimitReaderRead2(t *testing.T) {
	str := LimitReaderRead2()
	if str != "01234" {
		t.Errorf("TestLimitReaderRead Error %s" , str)
	}
}

func TestSectionReaderRead(t *testing.T) {
	str := SectionReaderRead()
	if str != "345" {
		t.Errorf("TestSectionReaderRead Error %s", str)
	}
}

func TestSectionReaderSeek(t *testing.T) {
	str := SectionReaderSeek()
	if str != "567" {
		t.Errorf("TestSectionReaderSeek Error %s", str)
	}
}

func TestSectionReaderReadAt(t *testing.T) {
	str := SectionReaderReadAt()
	if str != "789" {
		t.Errorf("TestSectionReaderReadAt Error %s", str)
	}
}

func TestSectionReaderSize(t *testing.T) {
	size := SectionReaderSize()
	if size != 9 {
		t.Errorf("TestSectionReaderSize Error %d", size)
	}
}

func TestTeeReaderRead(t *testing.T) {
	str := TeeReaderRead()
	if str != "012" {
		t.Errorf("TestTeeReaderRead Error %s", str)
	}
}