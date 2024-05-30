package main

import (
	"bufio"
	"os"
	"testing"
)

func Test1(t *testing.T) {
	os.Args = []string{"cmd", "-u", "test.txt"}
	if err := Sort(); err != nil {
		t.Fatal(err)
	}

	// читаем полученный файл
	file1, _ := os.Open("sorted.txt")
	scanner1 := bufio.NewScanner(file1)
	data1 := make([]string, 0)
	for scanner1.Scan() {
		line := scanner1.Text()
		data1 = append(data1, line)
	}

	// читаем  файл с ожидаемым результатом
	file2, _ := os.Open("expect.txt")
	scanner := bufio.NewScanner(file2)
	data2 := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		data2 = append(data2, line)
	}
	// выполняем сравнение
	for i := 0; i < len(data1); i++ {
		if data1[i] != data2[i] {
			t.Errorf("Not equal")
		}
	}

}
