package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры):
на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки
*/

type Sflags struct {
	N, R, U  bool
	K        int
	Filename string
}

func parseFlags(f *Sflags) {
	// flag.BoolVar(&f.K, "k", false, "указание колонки для сортировки")
	flag.BoolVar(&f.N, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&f.R, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&f.U, "u", false, "не выводить повторяющиеся строки")
	flag.IntVar(&f.K, "k", -1, "указание колонки для сортировки")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("No file specified")
		os.Exit(1)
	}
	for _, item := range args {
		f.Filename += item
	}
}

func openFile(f *Sflags) ([]string, error) {
	file, err := os.Open(f.Filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	data := make([]string, 0)
	uniqLines := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if f.U {
			if uniqLines[line] {
				continue
			} else {
				uniqLines[line] = true
			}
		} else {
			data = append(data, line)
		}
	}
	if f.U {
		for key := range uniqLines {
			data = append(data, key)
		}

	}
	return data, nil
}

func Reverse(s []string) []string {
	if len(s) <= 1 {
		return s
	}
	reverse := make([]string, 0)
	for i := len(s) - 1; i >= 0; i-- {
		reverse = append(reverse, s[i])
	}
	return reverse
}

func sortByColumn(s []string, f *Sflags) []string {
	column := f.K
	m := make(map[string]string)
	for _, item := range s {
		line := strings.Split(item, " ")
		if len(line)-1 >= column {
			m[line[column]] = item
		} else {
			m[line[0]] = item
		}
	}
	dataKeys := make([]string, 0)
	for key := range m {
		dataKeys = append(dataKeys, key)
	}
	result := make([]string, 0)
	if f.N {
		numbers := make([]int, 0)
		for _, str := range dataKeys {
			num, err := strconv.Atoi(str)
			if err != nil {
				continue
			}
			numbers = append(numbers, num)
		}
		sort.Ints(numbers)
		for i := 0; i < len(numbers); i++ {
			result = append(result, m[strconv.Itoa(numbers[i])])
		}
	} else {
		sort.Strings(dataKeys)
		for i := 0; i < len(dataKeys); i++ {
			result = append(result, m[dataKeys[i]])
		}
	}
	return result
}

func Sort() []string {
	var f Sflags
	parseFlags(&f)
	s, _ := openFile(&f)
	if f.K >= 0 {
		s = sortByColumn(s, &f)
	}
	if !f.N && f.K == -1 {
		sort.Strings(s)
	}
	if f.R {
		s = Reverse(s)
	}
	return s
}

func main() {
	for _, item := range Sort() {
		fmt.Println(item)
	}
}
