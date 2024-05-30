package main

import (
	"fmt"
	"sort"
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
слово из множества.
Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.
*/

const alphabet = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"

type Character struct {
	Word   string
	Arr    [33]int
	Lenght int
}

func (c *Character) Setter(s string) {
	c.Word = s
	runes := []rune(s)
	c.Lenght = len(runes)
	runess := []rune(alphabet)
	i := 0
	for _, r := range runes {
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

func (c *Character) Equal(other *Character) bool {
	return c.Lenght == other.Lenght && c.Arr == other.Arr
}

func

func CreateListCharacter(arr *[]string) []Character {
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

func FindAnagram(arr *[]string) /* map[string]*[]string */ {
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
		for i, w := range item {
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
	}

	for key, value := range anagram {
		fmt.Println(key + ":")
		if *value != nil {
			for _, word := range *value {
				fmt.Println("  - " + word)
			}
		} else {
			fmt.Println("  - No words found")
		}
	}
}

func main() {
	// a := []string{"кот", "ток", "карета", "ракета", "мама"}
	a := []string{"листок", "пятак", "слиток", "пятка", "столик", "тяпка", "тяпка"}
	FindAnagram(&a)
}
