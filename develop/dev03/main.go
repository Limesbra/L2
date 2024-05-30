package main

import (
	"bufio"
	"bytes"
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

// Структура которая содержит информацию о флагах и имени файла
type Sflags struct {
	N, R, U  bool
	K        int
	Filename string
}

// Функция парсинга. Парсит флаги и имя файла который нужно отсортировать
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

// Открываем файл и считываем всю информацию из него
// Если сортировка вызвана с флагом -u происходит удаление дубликатов
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

// функция обратной сортировки
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

// Функция сортировки по колонкам
// Разбивает строку на колонки и сортирует по номеру колонки
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

// Создаем новый файл и записываем в него результат сортировки
func writeFile(s []string) error {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			buf.WriteString(s[i])
			continue
		}
		buf.WriteString(s[i] + "\n")
	}
	file, err := os.Create("sorted.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(buf.Bytes())
	return err
}

// функция сортировки
func Sort() error {
	var f Sflags
	parseFlags(&f)
	s, _ := openFile(&f)
	if f.K >= 0 {
		s = sortByColumn(s, &f)
	}
	if f.N {
		f.K = 0
		s = sortByColumn(s, &f)
	}
	if !f.N && f.K == -1 {
		sort.Strings(s)
	}
	if f.R {
		s = Reverse(s)
	}
	err := writeFile(s)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {
	Sort()
}
