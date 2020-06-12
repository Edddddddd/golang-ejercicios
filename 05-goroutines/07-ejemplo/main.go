package main

import "fmt"

func printHello() {
	fmt.Println("Hello World!")
}

func main() {
	fmt.Println("main execution started")

	// create goroutine
	go printHello()

	fmt.Println("main execution stopped")
}
