package bytes

import (
	"testing"
)

func TestLen(t *testing.T) {
	if Len() != 10 {
		t.Errorf("TestLen Error. %d", Len())
	}
}

func TestSize(t *testing.T) {
	if Size() != 10 {
		t.Errorf("TestSize Error. %d", Size())
	}
}

func TestRead(t *testing.T) {
	str := Read()
	if str != "012" {
		t.Errorf("TestRead Error. %s", str)
	}
}

func TestReadAt(t *testing.T) {
	str := ReadAt()
	if str != "234" {
		t.Errorf("TestReadAt Error. %s", str)
	}
}

func TestReadByte(t *testing.T) {
	b := ReadByte()
	if b != 48 {
		t.Errorf("TestReadByte Error. %v", b)
	}
}

func TestUnreadByte(t *testing.T) {
	b := ReadByte()
	if b != 48 {
		t.Errorf("TestReadByte Error. %v", b)
	}
}

func TestReadRune(t *testing.T) {
	ch, size := ReadRune()
	if ch != 48 {
		t.Errorf("rune. %v", 48)
	}
	if size != 1 {
		t.Errorf("rune. %v", size)
	}
}

func TestUnreadRune(t *testing.T) {
	ch, size := UnreadRune()
	if ch != 48 {
		t.Errorf("rune. %v", 48)
	}
	if size != 1 {
		t.Errorf("rune. %v", size)
	}
}

func TestSeek(t *testing.T) {
	str := Seek()
	if str != "" {
		t.Errorf("TestSeek Error. %v", str)
	}
}

func TestSeek2(t *testing.T) {
	str := Seek2()
	if str != "0123456789" {
		t.Errorf("TestSeek Error. %v", str)
	}
}

func TestWriteTo(t *testing.T) {
	i := WriteTo()
	if i != 10 {
		t.Errorf("TestWriteTo Error. %v", i)
	}
}

func TestReset(t *testing.T) {
	str := Reset()
	if str != "" {
		t.Errorf("TestReset Error. %s", str)
	}
}

func TestReset2(t *testing.T) {
	str := Reset2()
	if str != "0123456789" {
		t.Errorf("TestReset Error. %s", str)
	}
}