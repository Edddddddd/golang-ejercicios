package main

import "fmt"

func main() {
	roc := make(<-chan int)
	soc := make(chan<- int)

	fmt.Printf("Canal de solo recepcion `%T`\n", roc)
	fmt.Printf("Canal de solo envio `%T\n", soc)
}
