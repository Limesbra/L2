package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Builder interface {
	BuildPart1(part1 int)
	BuildPart2(part2 string)
	GetResult() *ComplexObject
}

type ConcreteBuilder struct {
	Part1 int
	Part2 string
}

func (cb *ConcreteBuilder) BuildPart1(part1 int) {
	cb.Part1 = part1
}

func (cb *ConcreteBuilder) BuildPart2(part2 string) {
	cb.Part2 = part2
}

func (cb *ConcreteBuilder) GetResult() *ComplexObject {
	return &ComplexObject{Part1: cb.Part1, Part2: cb.Part2}
}

func GetBuilder() Builder {
	return &ConcreteBuilder{}
}

type ComplexObject struct {
	Part1 int
	Part2 string
}

func (c *ComplexObject) Print() {
	fmt.Println(c.Part1, c.Part2)
}
