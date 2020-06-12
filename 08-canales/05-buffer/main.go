package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func main() {
	fmt.Println("main() inicia")
	c := make(chan int, 3)

	go squares(c)

	c <- 1
	c <- 2
	c <- 3
	//c <- 4 // desbordando el buffer

	// close(c)// aun podemos ver estos datos
	for elem := range c {
		fmt.Println(elem)
	}

	fmt.Printf("Largo %v and capacidad  <%v", len(c), cap(c))
	fmt.Println("main() termina")
}
