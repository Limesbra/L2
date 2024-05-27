package main

import "testing"

func Test1(t *testing.T) {
	actual, _ := unpacker("a4bc2d5e")
	expected := "aaaabccddddde"
	if actual != expected {
		t.Errorf("actual %q, expected %q", actual, expected)
	}
}

func Test2(t *testing.T) {
	actual, _ := unpacker("")
	expected := ""
	if actual != expected {
		t.Errorf("actual %q, expected %q", actual, expected)
	}
}

func Test3(t *testing.T) {
	actual, _ := unpacker(" ")
	expected := " "
	if actual != expected {
		t.Errorf("actual %q, expected %q", actual, expected)
	}
}

func Test4(t *testing.T) {
	_, err := unpacker("45")
	if err == nil {
		t.Errorf("Should be an error")
	}
}
