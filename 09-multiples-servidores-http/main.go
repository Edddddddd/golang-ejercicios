package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1) // use only 1 processor core
	fmt.Println("Servidor iniciado....")
	// create a default route handler
	http.HandleFunc( "/",
		func( res http.ResponseWriter, req *http.Request ) {
			fmt.Fprint( res, "Hello: " + req.Host )
	} )

	// create a goroutine
	go func() {
		// spawn an HTTP server in `other` goroutine
		log.Fatal( http.ListenAndServe( ":9000", nil ) )
	}()

	// spawn an HTTP server in `main` goroutine
	log.Fatal( http.ListenAndServe( ":9001", nil ) )

}