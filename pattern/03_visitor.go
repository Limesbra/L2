package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// представляет собой интерфейс посетителя
type Visitor interface {
	VisitBar(b *Bar) string
	VisitBeach(beach *Beach) string
	VisitSchool(s *School) string
}

// интерфейс который принимает визит от посетителя
type Place interface {
	Accept(v Visitor) string
}

// реализует интерфейс  place
type Bar struct{}

// реализует функции Accept
func (b *Bar) Accept(v Visitor) string {
	return v.VisitBar(b)
}

// метод структуры Bar
func (b *Bar) IntoBar() string {
	return "drinks some dirnks"
}

// реализует интерфейс  place
type Beach struct{}

// реализует функции Accept
func (beach *Beach) Accept(v Visitor) string {
	return v.VisitBeach(beach)
}

// метод структуры Beach
func (beach *Beach) OnTheBeach() string {
	return "swimming"
}

// реализует интерфейс  place
type School struct{}

// реализует функции Accept
func (s *School) Accept(v Visitor) string {
	return v.VisitSchool(s)
}

// метод структуры School
func (s *School) AtSchool() string {
	return "be bored"
}

// реализует интерфейс Visitor
type ConcreteVisitor struct{}

// реализация метода VisitBar
func (c *ConcreteVisitor) VisitBar(b *Bar) string {
	return b.IntoBar()
}

// реализация метода VisitBeach
func (c *ConcreteVisitor) VisitBeach(beach *Beach) string {
	return beach.OnTheBeach()
}

// реализация метода VisitSchool
func (c *ConcreteVisitor) VisitSchool(s *School) string {
	return s.AtSchool()
}
