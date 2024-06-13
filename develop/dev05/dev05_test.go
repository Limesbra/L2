package main

import (
	"os"
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	except := map[int]string{0: "1:computer 1", 1: "2:mouse 2", 2: "3:LAPTOP 3", 3: "4:RedHat", 4: "5:data 4", 5: "6:RedHat 4 5 6", 6: "7:laptop 8",
		7: "8:debian 10", 8: "9:laptop 5", 9: "10:123ght 9"}
	os.Args = []string{"cmd", "-n", " ", "test.txt"}
	file := Grep()
	status := reflect.DeepEqual(except, *file)
	if !status {
		t.Errorf("Not equal")
	}
}

func Test2(t *testing.T) {
	except := map[int]string{0: "computer 1", 1: "mouse 2", 2: "LAPTOP 3", 3: "RedHat", 4: "data 4", 5: "RedHat 4 5 6",
		6: "debian 10", 7: "123ght 9"}
	os.Args = []string{"cmd", "-v", "laptop", "test.txt"}
	file := Grep()
	status := reflect.DeepEqual(except, *file)
	if !status {
		t.Errorf("Not equal")
	}
}

func Test3(t *testing.T) {
	except := map[int]string{3: "RedHat", 5: "RedHat 4 5 6"}
	os.Args = []string{"cmd", "-i", "redhat", "test.txt"}
	file := Grep()
	status := reflect.DeepEqual(except, *file)
	if !status {
		t.Errorf("Not equal")
	}
}

func Test4(t *testing.T) {
	except := "2"
	os.Args = []string{"cmd", "-i", "-c", "redhat", "test.txt"}
	file := Grep()
	// fmt.Println(*file)
	// fmt.Println(except)
	if (*file)[-1] != except {
		t.Errorf("Not equal")
	}
}
