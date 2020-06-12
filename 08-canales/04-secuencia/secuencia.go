package main

import "fmt"

func squaress(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}

	close(c) // close channel
}

func main() {
	fmt.Println("main() inicio")
	c := make(chan int)

	go squaress(c)

	for {
		val, ok := <-c
		if ok == false {
			fmt.Println(val, ok, "<-- Termina loop!")
			break
		} else {
			fmt.Println(val, ok)
		}
	}

	fmt.Println("main() Termina")
}
