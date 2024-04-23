package main

import (
	"book_test/db"
	"book_test/repositories"
	"book_test/routes"
	"book_test/services"
	"os"

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

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userRoutes := routes.NewUserRoutes(userService)

	bookRoutes.GetRoutes(e)
	userRoutes.GetRoutes(e)

	e.Logger.Fatal(e.Start(":" + envPortOr("8080")))
}

func envPortOr(port string) string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}
