package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main()  {
	fmt.Println("Iniciando Servidor apside")

	http.HandleFunc("/",
		func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprint(writer,"Hola Apside:"+ request.Host)
		})
	go func() {
		log.Fatal(http.ListenAndServe(":9000", nil))
	}()

	log.Fatal(http.ListenAndServe(":9001", nil))
}
