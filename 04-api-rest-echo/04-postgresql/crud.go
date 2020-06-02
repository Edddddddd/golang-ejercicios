package main

import (
	"fmt"
)

func main() {
	/*u:= Estudiantes {
		Nombre :"Gabriela Catalina Silva Gonzalez",
		Edad :37,
		Active: true,
	}
	err := Crear(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Creado Exitosamente")*/
	es, err := Consultar()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(es)
	a := 0
	fmt.Scan(&a)
}
