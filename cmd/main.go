package main

import (
	"book_test/db"
	"book_test/repositories"
	"book_test/routes"
	"book_test/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := db.InitDB()

	bookRepository := repositories.NewBookRepository(db)
	bookService := services.NewBookService(bookRepository)
	bookRoutes := routes.NewBookRoutes(bookService)

	bookRoutes.GetRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
