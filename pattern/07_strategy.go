package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type Strategy interface {
	Execute(n int)
}

type ConcreteStrategyA struct{}

func (a *ConcreteStrategyA) Execute(n int) {
	fmt.Println(n * n)
}

type ConcreteStrategyB struct{}

func (b *ConcreteStrategyB) Execute(n int) {
	fmt.Println(n * n * n)
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

func (c *Context) DoStrategy(n int) {
	c.strategy.Execute(n)
}
