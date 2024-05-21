package pattern

import "fmt"

// /*
// 	Реализовать паттерн «фасад».
// Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
// 	https://en.wikipedia.org/wiki/Facade_pattern
// */

// структура - фасад
type Band struct {
	V Vocalist
	D Drummer
	G Guitarist
	B Bassist
}

// конструктор
func NewBand() Band {
	return Band{V: Vocalist{},
		D: Drummer{},
		G: Guitarist{},
		B: Bassist{}}
}

// функция запускающая подсистемы
func (band *Band) MakeHit() {
	band.V.Sing()
	band.D.Play()
	band.G.Play()
	band.B.Play()
}

// структура - подсистема1
type Vocalist struct{}

// ультра сложная внутрянка подсистемы 1
func (v *Vocalist) Sing() {
	fmt.Println("Вокалист поет")
}

// структура - подсистема2
type Drummer struct{}

// ультра сложная внутрянка подсистемы 2
func (d *Drummer) Play() {
	fmt.Println("Барабанщик играет свою партию")
}

// структура - подсистема3
type Guitarist struct{}

// ультра сложная внутрянка подсистемы 3
func (g *Guitarist) Play() {
	fmt.Println("Гитарист играет свою партию")
}

// структура - подсистема4
type Bassist struct{}

// ультра сложная внутрянка подсистемы 4
func (b *Bassist) Play() {
	fmt.Println("Басист - самый крутой")
}
