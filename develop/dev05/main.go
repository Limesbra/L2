package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
)

/*
Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).

Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения +
-B - "before" печатать +N строк до совпадения +
-C - "context" (A+B) печатать ±N строк вокруг совпадения +
-c - "count" (количество строк) +
-i - "ignore-case" (игнорировать регистр) +
-v - "invert" (вместо совпадения, исключать) +
-F - "fixed", точное совпадение со строкой, не паттерн +
-n - "line num", напечатать номер строки +
*/
type Gflags struct {
	Count, Ignore, Invert, Fixed, Num, Border    bool
	After, Before, Context, BotBorder, TopBorder int
	Filename, Fstr                               string
	R                                            *regexp.Regexp
}

func (g *Gflags) setBorders() {
	if g.After == 0 && g.Before == 0 && g.Context == 0 {
		return
	} else {
		if g.After > 0 {
			g.BotBorder = g.After
		}
		if g.Before > 0 {
			g.TopBorder = g.Before
		}
		if g.After == 0 && g.Before == 0 {
			g.BotBorder = g.Context
			g.TopBorder = g.Context
		}
		g.Border = true
	}
}

func parseFlags(g *Gflags) {

	flag.BoolVar(&g.Count, "c", false, "количество строк")
	flag.BoolVar(&g.Ignore, "i", false, "игнорировать регистр")
	flag.BoolVar(&g.Invert, "v", false, "вместо совпадения, исключать")
	flag.BoolVar(&g.Fixed, "F", false, "точное совпадение со строкой, не паттерн")
	flag.BoolVar(&g.Num, "n", false, "напечатать номер строки")
	flag.IntVar(&g.After, "A", 0, "печатать +N строк после совпадения")
	flag.IntVar(&g.Before, "B", 0, "печатать +N строк до совпадения")
	flag.IntVar(&g.Context, "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("No arguments")
		os.Exit(1)
	}
	// если есть флаг -i
	if g.Ignore {
		g.R = regexp.MustCompile("(?i)" + args[0])
	} else {
		g.R = regexp.MustCompile(args[0])
	}
	if g.Fixed {
		g.Fstr = args[0]
	}
	g.Filename = args[1]
	// проверка на флаги -ABC и установка границ
	g.setBorders()

}

// Открываем файл и считываем всю информацию из него
func openFile(g *Gflags) (map[int]string, error) {
	file, err := os.Open(g.Filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	data := make(map[int]string)
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		data[i] = line
		i++
	}
	return data, nil
}

func borders(s *map[int]string, g *Gflags) map[int]string {
	out := make(map[int]string, 0)
	for i, str := range *s {
		if g.R.MatchString(str) && !g.Invert || !g.R.MatchString(str) && g.Invert {
			out[i] = str
		}
		if g.R.MatchString(str) {
			out[i] = str
			for j := g.TopBorder; j > 0; j-- {
				if i-j >= 0 {
					out[i-j] = (*s)[i-j]
				}
			}
			for k := g.BotBorder; k > 0; k-- {
				if i+k < len(*s) {
					out[i+k] = (*s)[i+k]
				}
			}
		}
	}
	return out
}

func compareBytes(s *map[int]string, g *Gflags) map[int]string {
	out := make(map[int]string, 0)
	sb := []byte(g.Fstr)
	for j, i := 0, 0; i < len(*s); i++ {
		fb := []byte((*s)[i])
		if !reflect.DeepEqual(sb, fb) {
			out[j] = (*s)[i]
			j++
		}
	}
	return out
}

func findInverse(s *map[int]string, g *Gflags) map[int]string {
	out := make(map[int]string, 0)
	if !g.Fixed {
		for j, i := 0, 0; i < len(*s); i++ {
			if !g.R.MatchString((*s)[i]) {
				out[j] = (*s)[i]
				j++
			}
		}
	} else {
		out = compareBytes(s, g)
	}
	return out
}

func addNumberRow(s *map[int]string) {
	for i := 0; i < len(*s); i++ {
		(*s)[i] = strconv.Itoa(i+1) + ":" + (*s)[i]
	}
}

func printResult(s *map[int]string) {
	for i, j := 0, 0; j < len(*s); i++ {
		if (*s)[i] != "" {
			j++
			fmt.Println((*s)[i])
		}
	}
}

func Grep() {
	var g Gflags
	parseFlags(&g)
	file, _ := openFile(&g)

	var output map[int]string

	// если есть флаг -n
	if g.Num && !g.Count {
		addNumberRow(&file)
	}
	// если есть флаг -v
	if g.Invert && !g.Border {
		output = findInverse(&file, &g)
	} else if g.Border { // если есть флаги -АВС
		output = borders(&file, &g)
	} else if g.Fixed { // если тольк есть флаг -F
		output = compareBytes(&file, &g)
	} else { // если обычный греп
		output = file
	}
	// если есть флаг -c
	if g.Count {
		fmt.Println(len(output))
		return
	}
	printResult(&output)
}

func main() {
	Grep()
}
