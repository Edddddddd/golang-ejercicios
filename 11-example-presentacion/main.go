// Las _Interfaces_ son colecciones nombradas
// de signaturas de métodos.

package main

import "fmt"
import "math"

type geometrica interface {
	area() float64
	perim() float64
}

type cuadro struct {
	ancho, altura float64
}
type circulo struct {
	radio float64
}

func (s cuadro) area() float64 {
	return s.ancho * s.altura
}
func (s cuadro) perim() float64 {
	return 2*s.ancho + 2*s.altura
}

func (c circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}
func (c circulo) perim() float64 {
	return 2 * math.Pi * c.radio
}

func medida(g geometrica) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	s := cuadro{ancho: 3, altura: 4}
	c := circulo{radio: 5}

	medida(s)
	medida(c)
}
