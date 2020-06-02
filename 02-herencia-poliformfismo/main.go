package main

import (
	us "./user"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var texto string
var status bool = true
var numero1 int
var resultado int
var leyenda string

// Herencia
type pepe struct {
	us.Usuario
	// Puedo darle polimorfismo

}

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	numero2, numero3, numero4, texto, status := 2, 5, 67, "Hola", false
	var moneda int64 = 0
	numero2 = int(moneda)
	texto = fmt.Sprintf("%d", moneda)

	texto = strconv.Itoa(int(moneda))
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
	fmt.Println(status)
	status = true
	if status = false; status == true {
		fmt.Println("Es verdad")
	} else {
		fmt.Println("Es falso")
	}

	switch numero12 := 3; numero12 {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		fmt.Println("Mayor a 3")
	}

	fmt.Println("Ingrese numero 1: ")
	fmt.Scanln(&numero1)
	fmt.Println("El numero ingresado es", numero1)

	fmt.Println("Ingrese numero 2: ")
	fmt.Scanln(&numero)
	fmt.Println("El numero ingresado es", numero)
	fmt.Println("Multiplicacion", numero1+numero)

	fmt.Println("Ingresa la Leyenda")
	scaner := bufio.NewScanner(os.Stdin)
	if scaner.Scan() {
		leyenda = scaner.Text()
	}

	resultado = numero1 + numero
	fmt.Println(leyenda, resultado)
	i := 0

	//FOR:
	count := 10
	for i < count {
		println(i)
		if i == 3 {
			i++
			println("salgo a GoTo")
			//			goto FOR
		}

		if i == 4 {
			i = 5
			continue
		}
		if i == 5 {
			break
		}
		i++
	}

	valor, estado := uno(1)
	fmt.Println(valor)
	fmt.Println(estado)

	fmt.Printf("Sumo 5+7 = %d \n", Calculo(5, 7))

}

func uno(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

// Funciones anonimas
/****
Funciones que no tienen nombre
modificar funciones en tiempo de ejecucion

**/
