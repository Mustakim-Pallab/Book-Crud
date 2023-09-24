package routes

import (
	"book-crud/controller"
	"github.com/labstack/echo/v4"
)

func BookRoutes(e *echo.Echo) {
	book := e.Group("/bookstore")

	book.POST("/book", controller.CreateBook)
	book.GET("/book", controller.GetBooks)
	book.PUT("/book/:bookID", controller.UpdateBook)
	book.DELETE("/book/:bookID", controller.DeleteBook)
}
