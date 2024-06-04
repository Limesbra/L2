package main

import (
	"flag"
	"fmt"
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
	f int
	d string
	s bool
}

func parseFlags(c *cFlags) []string {
	flag.IntVar(&c.f, "f", 0, "выбрать поля (колонки)")
	flag.StringVar(&c.d, "d", "\t", "только строки с разделителем")
	flag.BoolVar(&c.s, "s", false, "только строки с разделителем")
	flag.Parse()
	// for i := 0; i < len(os.Args); i++ {
	// fmt.Println(os.Args[i], i)
	// }
	args := flag.Args()
	// fmt.Println(args[0])
	return args
}

// func worker(s *bufio.Scanner) string {
// 	var res strings.Builder
// 	for s.Scan() {
// 		data := s.Text()
// 		res.WriteString(data + " ")
// 	}

// 	r := res.String()

// 	return r
// }

func worker(str []string) string {
	var res strings.Builder
	for _, s := range str { // Используем s.Scan() вместо (*s).Scan(), так как s уже указывает на Scanner
		res.WriteString(s + " ") // Убедитесь, что добавляете пробел между элементами, если это необходимо
	}

	r := res.String()
	return r
}

func Cut() {
	var c cFlags
	expression := parseFlags(&c)
	fmt.Println(worker(expression))

}

func main() {
	Cut()
}
