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

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)


В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.
*/

func unpacker(s string) (string, error) {

	runes := []rune(s)
	if unicode.IsDigit(runes[0]) {
		return "", fmt.Errorf("invalid string")
	}

	var unpackRunes []rune
	var oldRune rune
	for i, r := range runes {
		if i%2 != 0 && unicode.IsDigit(r) {
			num, err := strconv.Atoi(string(r))
			if err != nil {
				return "", err
			}
			for j := 0; j <= num; j++ {
				unpackRunes = append(unpackRunes, oldRune)
			}
		}
		oldRune = r
	}
	str := string(unpackRunes)
	return str, nil
}

func main() {
	s := "a4"
	s1, _ := unpacker(s)
	fmt.Println(s1)
}
