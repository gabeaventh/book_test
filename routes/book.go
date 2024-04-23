package routes

import (
	"book_test/models"
	"book_test/services"
	"book_test/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookRoutes struct {
	bookService services.BookService
	userService services.UserService
}

func NewBookRoutes(bookService services.BookService, userService services.UserService) *BookRoutes {
	return &BookRoutes{
		bookService: bookService,
		userService: userService,
	}
}

func (r *BookRoutes) GetAllBooks(c echo.Context) error {
	books, err := r.bookService.GetAllBooks(c)
	if err != nil {
		return handleError(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, "Books fetched successfully", books)
}

func (r *BookRoutes) CreateBook(c echo.Context) error {
	auth := r.GetUser(c)

	if auth == nil {
		return auth
	}

	book := &models.Book{}
	if err := c.Bind(book); err != nil {
		return handleError(c, http.StatusBadRequest, err.Error())
	}

	if err := validateBookInput(book); err != nil {
		return handleError(c, http.StatusBadRequest, err.Error())
	}

	book, err := r.bookService.CreateBook(c, book)
	if err != nil {
		return err
	}
	return utils.SuccessResponse(c, "Book created successfully", book)
}

func (r *BookRoutes) UpdateBook(c echo.Context) error {
	auth := r.GetUser(c)
	if auth == nil {
		return auth
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return handleError(c, http.StatusBadRequest, "Invalid ID")
	}
	book := &models.Book{}
	if err := c.Bind(book); err != nil {

		return handleError(c, http.StatusBadRequest, "Invalid request body")
	}

	if err := validateBookInput(book); err != nil {
		return handleError(c, http.StatusBadRequest, err.Error())
	}

	book.ID = idInt
	book, err = r.bookService.UpdateBook(c, book)
	if err != nil {
		return err
	}
	return utils.SuccessResponse(c, "Book updated successfully", book)
}

func (r *BookRoutes) DeleteBook(c echo.Context) error {
	auth := r.GetUser(c)

	if auth == nil {
		return auth
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return handleError(c, http.StatusBadRequest, "Invalid ID")
	}

	err = r.bookService.DeleteBook(c, idInt)
	if err != nil {
		return err
	}
	return utils.SuccessResponse(c, "Book deleted successfully", nil)
}

func (r *BookRoutes) GetBookByID(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return handleError(c, http.StatusBadRequest, "Invalid ID")
	}

	if idInt == 0 {
		return handleError(c, http.StatusBadRequest, "ID is required")
	}
	book, err := r.bookService.GetBookByID(c, idInt)
	if err != nil {
		return err
	}
	return utils.SuccessResponse(c, "Book fetched successfully", book)
}

func (r *BookRoutes) GetUser(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")

	if token == "" {
		return handleError(c, http.StatusUnauthorized, "Unauthorized")
	}

	user, err := r.userService.GetUser(token, c)

	if err != nil {
		return err
	}

	return utils.SuccessResponse(c, "Authorized", user)
}

func handleError(c echo.Context, statusCode int, message string) error {
	return utils.ErrorResponse(c, statusCode, message)
}

func validateBookInput(book *models.Book) error {
	if book.Title == "" || book.Author == "" {
		return fmt.Errorf("title and author are required")
	}
	return nil
}

func (r *BookRoutes) GetRoutes(e *echo.Echo) {
	group := e.Group("/book")

	group.GET("", r.GetAllBooks)
	group.POST("", r.CreateBook)
	group.GET("/:id", r.GetBookByID)
	group.PUT("/:id", r.UpdateBook)
	group.DELETE("/:id", r.DeleteBook)
}
