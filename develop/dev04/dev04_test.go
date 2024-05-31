package main

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	expect := map[string][]string{
		"ток":    {"кот", "ток"},
		"карета": {"карета", "ракета"},
	}
	a := []string{"ток", "кот", "карета", "РАКЕТА", "мама"}
	res := FindAnagram(&a)

	for key := range expect {
		if !reflect.DeepEqual(expect[key], *res[key]) {
			t.Errorf("Not equal")
		}
	}

}
