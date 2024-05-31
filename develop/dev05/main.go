package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

/*
Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).

Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк) +
-i - "ignore-case" (игнорировать регистр) +
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки +
*/
type Gflags struct {
	Count, Ignore, Invert, Fixed, Num, Border    bool
	After, Before, Context, BotBorder, TopBorder int
	Filename                                     string
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
	g.Filename = args[1]
	// проверка на флаги -ABC и установка границ
	g.setBorders()

}

// Открываем файл и считываем всю информацию из него
func openFile(g *Gflags) ([]string, error) {
	file, err := os.Open(g.Filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	data := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
	}
	return data, nil
}

// // Функция копирует из слайса в мапу данные из заданного интервала в слайсе
// func copyToMapByInterval(m map[int]string, data []string, first int, last int) {
// 	for k := first; k < last; k++ {
// 		m[k] = data[k]
// 	}
// }

// // Реализует ключ -B
// func grepBefore(g *Gflags, s *[]string) map[int]string {
// 	buf := make(map[int]string)
// 	for i, str := range *s {
// 		if g.R.MatchString(str) && !g.Invert || !g.R.MatchString(str) && g.Invert {
// 			buf[i] = str
// 			if i-g.TopBorder >= 0 {
// 				copyToMapByInterval(buf, *s, i-g.TopBorder, i)
// 			} else {
// 				copyToMapByInterval(buf, *s, 0, i)
// 			}
// 		}
// 	}
// 	return buf
// }

// // Реализует ключ -A
// func grepAfter(g *Gflags, s *[]string) map[int]string {
// 	buf := make(map[int]string)
// 	for i, str := range *s {
// 		if g.R.MatchString(str) && !g.Invert || !g.R.MatchString(str) && g.Invert {
// 			buf[i] = str
// 			if i+g.BotBorder < len(*s) {
// 				copyToMapByInterval(buf, *s, i+1, i+g.BotBorder+1)
// 			} else {
// 				copyToMapByInterval(buf, *s, i+1, len(*s))
// 			}
// 		}
// 	}
// 	return buf
// }

// // Реализует ключ -C
// func grepContext(g *Gflags, s *[]string) map[int]string {
// 	buf := make(map[int]string)
// 	for i, str := range *s {
// 		if g.R.MatchString(str) && !g.Invert || !g.R.MatchString(str) && g.Invert {
// 			buf[i] = str

// 			if i-g.Context >= 0 {
// 				copyToMapByInterval(buf, *s, i-g.TopBorder, i)
// 			} else {
// 				copyToMapByInterval(buf, *s, 0, i)
// 			}

// 			if i+g.Context < len(*s) {
// 				copyToMapByInterval(buf, *s, i+1, i+g.TopBorder+1)
// 			} else {
// 				copyToMapByInterval(buf, *s, i+1, len(*s))
// 			}
// 		}
// 	}

// 	return buf
// }

func FindInverse(s *[]string, g *Gflags) []string {
	output := make([]string, 0)
	if g.Border { // нужно сделать мапу
		// m:= make(map[string]bool)
		for i := 0; i < len(*s); i++ {
			if !g.R.MatchString((*s)[i]) {
				// m[(*s)[i]] = true
				output = append(output, (*s)[i])
			} else {
				for j := g.TopBorder; j > 0; j-- {
					output = append(output, (*s)[i-j])
				}

			}
		}
	}
	for _, item := range *s {
		if !g.R.MatchString(item) {
			output = append(output, item)
		}
	}
	return output
}

func AddNumberRow(s *[]string) []string {
	out := make([]string, len(*s))
	for i := 0; i < len(*s); i++ {
		out[i] = strconv.Itoa(i+1) + ":" + (*s)[i]
	}
	return out
}

func Grep() {
	var g Gflags
	parseFlags(&g)
	file, _ := openFile(&g)

	var output []string

	// если есть флаг -n
	if g.Num && !g.Count {
		output = AddNumberRow(&file)
	}
	// если есть флаг -v
	if g.Invert {
		output = FindInverse(&file, &g)
	}
	// если есть флаг -c
	if g.Count {
		fmt.Println(len(output))
		return
	}
	// if g.After
	fmt.Println(g.After)

	// for _, item := range file {
	// 	if g.R.MatchString(item) && !g.Invert {
	// 		output = append(output, item)
	// 	} else if !g.R.MatchString(item) && g.Invert {
	// 		output = append(output, item)
	// 	}
	// }

	// for _, s := range output {
	// 	fmt.Println(s)
	// }

}

func main() {
	Grep()
}
