package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// параметр помогает определить доступные действия
type parametr string

const (
	X parametr = "X"
	Y parametr = "Y"
	Z parametr = "Z"
)

// интерфейс фабрики
type Creator interface {
	Create(p parametr) Product
}

// интерфейс продукта
type Product interface {
	Do()
}

// реализация интерфейса фабрики
type Fabric struct{}

// реализация метода из интерфейса  Creator
func (f *Fabric) Create(p parametr) Product {
	var product Product
	switch p {
	case X:
		product = &RealProductA{parametr: string(X)}
	case Y:
		product = &RealProductY{parametr: string(Y)}
	case Z:
		product = &RealProductY{parametr: string(Z)}
	default:
		fmt.Println("Unknown parametr")
	}
	return product
}

// реализация интерфейса продукта
type RealProductA struct {
	parametr string
}

// реализация метода из интерфейса Product
func (a *RealProductA) Do() {
	fmt.Println(a.parametr)
}

// реализация интерфейса продукта
type RealProductY struct {
	parametr string
}

// реализация метода из интерфейса Product
func (y *RealProductY) Do() {
	fmt.Println(y.parametr)
}

// реализация интерфейса продукта
type RealProductZ struct {
	parametr string
}

// реализация метода из интерфейса Product
func (z *RealProductZ) Do() {
	fmt.Println(z.parametr)
}
