package main

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("Hello World!")
}

func main() {
	fmt.Println("main inicio ")

	// creo goroutine
	go printHello()

	// hago esperar el programa para ver la ejecucion de la goroutine
	time.Sleep(10 * time.Millisecond)
	fmt.Println("main finalizo")
}
