package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// интерфейс различных состояний
type State interface {
	Handle()
}

// контекст
type Conditions struct {
	state State
}

// функция показывающая текущее состояние
func (c *Conditions) MyState() {
	c.state.Handle()
}

// сеттер для установки состояния
func (c *Conditions) SetState(state State) {
	c.state = state
}

// состояние А, реализация интерфейса состояния
type StateA struct{}

// реализация метода интерфейса для StateA
func (a *StateA) Handle() {
	fmt.Println("Состояние A")
}

// состояние А, реализация интерфейса состояния
type StateB struct{}

// реализация метода интерфейса для StateB
func (b *StateB) Handle() {
	fmt.Println("Состояние B")
}
