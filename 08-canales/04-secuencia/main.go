package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}

	close(c) // cierro el canal
}

func main() {
	fmt.Println("main() inicia")
	c := make(chan int)

	go squares(c) // inicia goroutine

	// bloqueo y desbloqueo de la goroutine principal
	// hasta que cierre el canal
	for val := range c {
		fmt.Println(val)
	}

	fmt.Println("main() termina")
}
