package main

import "fmt"

type SerVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
}

type vegetal interface {
	Clasificacion() string
}

type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}

type mujer struct {
	hombre
	// Con solo el nombre se esta heredando todo lo del objeto hombre
}

func (h *hombre) estaVivo() { h.vivo = true }
func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) pensar()   { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre == true {
		return "Hombre"
	} else {
		return "Mujer"
	}
}

func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un %s, y ya estoy respirando \n", hu.sexo())
}

/*------------*/

type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (p *perro) estaVivo()         { p.vivo = true }
func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.comiendo = true }
func (p *perro) EsCarnivoro() bool { return p.carnivoro }

func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}

func AnimalesCarnivoro(an animal) int {
	if an.EsCarnivoro() == true {
		return 1
	}
	return 0
}

/*  */
func estoyVivo(v SerVivo) bool {
	return v.estaVivo()
}

func main() {
	totalCarnivoros := 0

	Dogo := new(perro)
	Dogo.carnivoro = true
	Dogo.vivo = true
	AnimalesRespirar(Dogo)
	totalCarnivoros = +AnimalesCarnivoro(Dogo)
	fmt.Printf("Total de carnivoros %d \n", totalCarnivoros)
	//fmt.Printf("Estoy vivo  =%d \n", totalCarnivoros)

	Pedro := new(hombre)
	Pedro.esHombre = true
	Pedro.vivo = true
	HumanosRespirando(Pedro)

	Maria := new(mujer)
	Maria.vivo = true
	HumanosRespirando(Maria)

}
