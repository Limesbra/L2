package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
	Паттерн Chain Of Responsibility относится к поведенческим паттернам уровня объекта.

Паттерн Chain Of Responsibility позволяет избежать привязки объекта-отправителя запроса к объекту-получателю запроса, при этом давая шанс обработать этот
запрос нескольким объектам. Получатели связываются в цепочку, и запрос передается по цепочке, пока не будет обработан каким-то объектом.

По сути это цепочка обработчиков, которые по очереди получают запрос, а затем решают, обрабатывать его или нет. Если запрос не обработан,
то он передается дальше по цепочке. Если же он обработан, то паттерн сам решает передавать его дальше или нет. Если запрос не обработан ни одним обработчиком,
то он просто теряется.


*/
// интерфейс обработчика
type Handler interface {
	Request(msg int) string
}

// реализация интерфейса Handler
type ConcreteHandlerA struct {
	next Handler
}

// реализация метода Request
func (a *ConcreteHandlerA) Request(msg int) string {
	if msg < 0 {
		return "less then zero"
	}
	return a.next.Request(msg)
}

// реализация интерфейса Handler
type ConcreteHandlerB struct {
	next Handler
}

// реализация метода Request
func (b *ConcreteHandlerB) Request(msg int) string {
	if msg == 0 {
		return "equel zero"
	}
	return b.next.Request(msg)
}

// реализация интерфейса Handler
type ConcreteHandlerC struct {
	next Handler
}

// реализация метода Request
func (c *ConcreteHandlerC) Request(msg int) string {
	if msg > 0 {
		return "greater zero"
	}
	return ""
}
