package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	u "github.com/eddy/echo-labs/03-services-respose/services"
	"os"
)

func main()  {

	e := echo.New()
	logger, err := os.OpenFile(
		"logs.log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
		)
	if err != nil {
		log.Fatal("no se pudo crear o abrir el archivo %v", err)
	}

	defer logger.Close()
	loggerConfig := middleware.LoggerConfig{
		Output: logger,
	}

	e.Use(middleware.LoggerWithConfig(loggerConfig))
	e.Static("/", "public")
	e.GET("/gopher/:name", gopherName)
	go e.GET("/services-gorilas/users/all",getUsuarios)
	go e.GET("/services-gorilas/users/all",getUsuarios)
	go e.GET("/services-gorilas/users/all",getUsuarios)
	go e.GET("/services-gorilas/users/all",getUsuarios)

	go func(s string) {
		fmt.Println(s)
	}("hola Apside")

	e.Start(":9080")
}

func gopherName(c echo.Context) error {
	p := c.Param("name")
	if p == "jpg" {
		return c.File("imgs/gopher.jpg")
	} else if p == "att" {
		return c.Attachment("imgs/gopher.jpg","gopher.jpg")
	}
	return c.HTML(http.StatusNotFound, "<h1>Error 404</h1>")
}

func getUsuarios(context echo.Context) error {
	es, err := u.ConsultarUsuarioGorilas()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, "Error")
	}
	return context.JSON(http.StatusOK, es )
}