package main

import (
	"fmt"
	"strconv"
	"unicode"
)

/*
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""


В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.
*/

func unpacker(s string) (string, error) {

	if len(s) == 0 {
		return "", nil
	}

	runes := []rune(s)
	if unicode.IsDigit(runes[0]) {
		return "", fmt.Errorf("invalid string")
	}

	res := make([]rune, 0)
	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			number := make([]rune, 0)
			for i < len(runes) && unicode.IsDigit(runes[i]) {
				number = append(number, runes[i])
				i++
			}
			i--
			num, _ := strconv.Atoi(string(number))
			for j := 0; j < num-1; j++ {
				res = append(res, res[len(res)-1])
			}
		} else {
			res = append(res, runes[i])
		}
	}
	return string(res), nil

}

func main() {
	s1, err := unpacker("* *")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s1)
}
