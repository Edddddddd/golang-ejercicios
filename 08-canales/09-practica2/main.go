
package main

import "fmt"

func main() {

	mensajes := make(chan string)

	go func() { mensajes <- "ping" }()

	msg := <-mensajes
	fmt.Println(msg)
}
