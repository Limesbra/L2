package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// интерфейс
type Command interface {
	Execute()
}

// реализует интерфейс Command
type Action1Command struct {
	r *Receiver
}

// реализация Execute()
func (a1c *Action1Command) Execute() {
	a1c.r.Action1()
}

// реализует интерфейс Command
type Action2Command struct {
	r *Receiver
}

// реализация Execute()
func (a2c *Action2Command) Execute() {
	a2c.r.Action2()
}

// реализует интерфейс Command
type Action3Command struct {
	r *Receiver
}

// реализация Execute()
func (a3c *Action3Command) Execute() {
	a3c.r.Action3()
}

// реализует инициатора, записывающего команду и провоцирующего её выполнение
type Invoker struct {
	commands []Command
}

// добавляем команду в очередь
func (i *Invoker) AddCommand(c Command) {
	i.commands = append(i.commands, c)
}

// удаляем команду в очередь
func (i *Invoker) DeleteCommand(c Command) {
	if len(i.commands) > 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

// выполняем команды в очереди
func (i *Invoker) Execute() {
	for _, item := range i.commands {
		item.Execute()
	}
}

// реализует получателя и имеет набор действий, которые команда можем запрашивать
type Receiver struct{}

func (r *Receiver) Action1() {
	fmt.Println(1)
}
func (r *Receiver) Action2() {
	fmt.Println(2)
}
func (r *Receiver) Action3() {
	fmt.Println(3)
}
