package interfaces

import "fmt"

//*************************************************
type GetterPrecio interface {
	GetPrecio()
}

//*************************************************
type Person1 struct {
	Precio float64
}

func NewPerson1(precio float64) (person1 Person1) {
	person1.Precio = precio
	return person1
}

func (p *Person1) GetPrecio() {
	p.Precio = 200.0
}

//*************************************************
type Person2 struct {
	Precio float64
}

func NewPerson2(precio float64) (person2 Person2) {
	person2.Precio = precio
	return person2
}

func (p *Person2) GetPrecio() {
	p.Precio = 100.00
}

//*************************************************
func GetterPrecioAll(gs ...GetterPrecio) {
	for _, v := range gs {
		v.GetPrecio()
	}
}

//*************************************************
func Interfaces() {
	var var1 float64 = 25.25
	var var2 float64 = 50.50
	p1 := NewPerson1(var1)
	p2 := NewPerson2(var2)
	var g1 GetterPrecio = &p1
	var g2 GetterPrecio = &p2
	fmt.Printf("%v\n", g1)
	fmt.Printf("%v\n", g2)
	GetterPrecioAll(g1, g2)
	fmt.Printf("%v\n", g1)
	fmt.Printf("%v\n", g2)
}
