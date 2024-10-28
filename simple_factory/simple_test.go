package simple_factory

import "testing"

func TestType(t *testing.T) {
	API := NewAPI(1)
	s := API.Say("Key")
	if s != "Hi, Key" {
		t.Fatal("Type1 test fail")
	}
}

func TestType2(t *testing.T) {
	API := NewAPI(2)
	s := API.Say("Key")
	if s != "Hello, Key" {
		t.Fatal("Type2 test fail")
	}
}
