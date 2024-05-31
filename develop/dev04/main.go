package main

import (
	"fmt"
	"sort"
	"strings"
	// "strings"
)

/*
Написать функцию поиска всех множеств анаграмм по словарю.

Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Требования:
Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
Выходные данные: ссылка на мапу множеств анаграмм
Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
слово из множества. +
Массив должен быть отсортирован по возрастанию. +
Множества из одного элемента не должны попасть в результат. +
Все слова должны быть приведены к нижнему регистру. +
В результате каждое слово должно встречаться только один раз. +
*/

const alphabet = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"

type Character struct {
	Word string
	Arr  [33]int
	// Lenght int
}

func (c *Character) Setter(s string) {
	c.Word = s
	runess := []rune(alphabet)
	i := 0
	for _, r := range s {
		for _, rr := range runess {
			if string(r) == string(rr) {
				c.Arr[i] += 1
				i = 0
				break
			}
			i++
		}
	}
}

func KillDoubling(arr *[]string) {
	m := make(map[string]bool)
	j := 0
	for _, item := range *arr {
		if !m[item] {
			m[item] = true
			(*arr)[j] = item
			j++
		}
	}
	*arr = (*arr)[:j]
}

func CreateListCharacter(arr *[]string) []Character {
	KillDoubling(arr)
	SortByLenght(arr)
	list := make([]Character, len(*arr))
	for i := 0; i < len(*arr); i++ {
		list[i].Setter((*arr)[i])
	}
	return list
}

func SortByLenght(arr *[]string) {
	sort.Slice(*arr, func(i, j int) bool {
		return len((*arr)[i]) < len((*arr)[j])
	})
}

func ToLower(arr *[]string) {
	for i, item := range *arr {
		(*arr)[i] = strings.ToLower(item)
	}
}

func FindAnagram(arr *[]string) map[string]*[]string {
	ToLower(arr)
	list := CreateListCharacter(arr)
	m := make(map[[33]int][]Character)
	for _, item := range list {
		if len(m) == 0 || m[item.Arr] == nil {
			a := make([]Character, 0)
			a = append(a, item)
			m[item.Arr] = a
		} else {
			m[item.Arr] = append(m[item.Arr], item)
		}
	}

	anagram := make(map[string]*[]string)
	for _, item := range m {
		key := item[0].Word
		flag := true
		for i, w := range item {
			if len(item) == 1 {
				flag = false
				continue
			}
			if i == 0 {
				a := make([]string, 0)
				a = append(a, w.Word)
				anagram[key] = &a
			} else {
				arr := anagram[key]
				*(arr) = append(*(arr), w.Word)
				anagram[key] = arr
			}
		}
		if flag {
			sort.Strings(*anagram[key])
		}
	}
	return anagram
}

func main() {
	a := []string{"ток", "кот", "карета", "РАКЕТА", "мама"}
	// a := []string{"листок", "пятак", "слиток", "пятка", "столик", "тяпка", "тяпка"}
	m := FindAnagram(&a)
	for key, value := range m {
		fmt.Println("Key:", key)
		fmt.Println("Value:", *value)
		fmt.Println()
	}
}
