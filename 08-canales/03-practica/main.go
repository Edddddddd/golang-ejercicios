package main

import "fmt"

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
}

func main() {
	fmt.Println("main() started")
	c := make(chan string, 1)

	go greet(c)
	c <- "John"
	c <- "Emiliano"

	fmt.Println("main() stopped")
}
