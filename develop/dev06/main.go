package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Реализовать утилиту аналог консольной команды cut (man cut).
Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/

type cFlags struct {
	d, f string
	s    bool
}

func parseFlags(c *cFlags) {
	flag.StringVar(&c.f, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&c.d, "d", "", "только строки с разделителем")
	flag.BoolVar(&c.s, "s", false, "только строки с разделителем")
	flag.Parse()
}

func worker(c *cFlags) error {
	column := make([]int, 0)

	switch {
	case strings.Contains(c.f, "-"): // если ввели диапазон пр: 1-5
		c := strings.Split(c.f, "-")
		start, _ := strconv.Atoi(c[0])
		end, _ := strconv.Atoi(c[1])
		for i := start; i <= end; i++ {
			column = append(column, i)
		}
	case strings.Contains(c.f, ","): // если ввели через , пр: 1,2,3
		c := strings.Split(c.f, ",")
		for _, item := range c {
			v, _ := strconv.Atoi(item)
			column = append(column, v)
		}
	default: // ввели одну цифру пр: 2
		v, _ := strconv.Atoi(c.f)
		column = append(column, v)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		str := scanner.Text()
		if c.d == "" && !c.s {
			fmt.Print(str)
		} else {
			s := strings.Split(str, c.d)
			if c.f == "" {
				return fmt.Errorf("usage: cut -f list [-s] [-w | -d delim] [file ...]")
			}
			for i, j := range column {
				if i == len(column)-1 {
					fmt.Print(s[0+j-1])
					break
				}
				fmt.Print(s[0+j-1], "\t")
			}
		}

	}
	return nil
}

func Cut() {
	var c cFlags
	parseFlags(&c)
	worker(&c)
}

func main() {
	Cut()
}
