package main

import (
	"book-crud/container"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	container.Serve(e)
}
