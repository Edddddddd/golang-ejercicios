package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Se utiliza para esperar que termine la ejecucion.
var wg sync.WaitGroup

func main()  {

	wg.Add(2)
	fmt.Println("Iniciando...")

	go imprimirCantidad("A")

	go imprimirCantidad("B")

	fmt.Println("Esperando que finalicen...")
	wg.Wait()
	fmt.Println("\n Terminando el programa")

}

func imprimirCantidad(etiqueta string)  {
	defer wg.Done()
	for cantidad := 1 ;cantidad <= 10; cantidad++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Microsecond)
		fmt.Printf("Cantidad: %d de %s\n", cantidad,etiqueta)
	}

}