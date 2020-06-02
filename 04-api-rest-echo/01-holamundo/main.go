package main

import (
	"github.com/labstack/echo"
	"net/http"
)

type Persona struct {
	FristName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hola Mundo")
	})
	e.GET("/json", func(c echo.Context) error {
		p := Persona{
			FristName: "Alexiys",
			LastName:  "Lazada",
			Age:       38,
		}
		return c.JSON(http.StatusOK, p)
	})
	e.Start(":8080")
}
