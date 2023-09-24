package container

import (
	"book-crud/config"
	"book-crud/connection"
	"book-crud/controller"
	"book-crud/repositories"
	"book-crud/routes"
	"book-crud/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

func Serve(e *echo.Echo) {
	config.InitConfig()

	db := connection.GetDB()
	bookRepo := repositories.BookDBInstance(db)
	bookService := service.BookServiceInstance(bookRepo)

	controller.SetBookService(bookService)
	routes.BookRoutes(e)
	log.Fatal(e.Start(fmt.Sprintf(":%s", config.InitConfig().Port)))
}
