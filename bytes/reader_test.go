package bytes

import "testing"

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