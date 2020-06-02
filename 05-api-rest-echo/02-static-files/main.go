package main

import "github.com/labstack/echo"

func main()  {
	e := echo.New()
	e.File("/", "public")
	e.Start(":8080")
}
