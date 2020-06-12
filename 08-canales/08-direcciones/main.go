package main

import "fmt"

// solo canal para enviar valores
func ping( pings chan <- string, mensaje string)  {
	pings <- mensaje
}

// funcion recibe valores y envia
func  pong(pings <-chan string, pongs chan<- string)  {
	// envia un canal recibido
	msg := <-pings
	// acepta un canal recibido
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "mensaje enviado")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}