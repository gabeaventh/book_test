package routes

import (
	"book_test/models"
	"book_test/services"
	"book_test/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookRoutes struct {
	bookService services.BookService
}

func NewBookRoutes(bookService services.BookService) *BookRoutes {
	return &BookRoutes{
		bookService: bookService,
	}
}

func (r *BookRoutes) GetAllBooks(c echo.Context) error {
	books, err := r.bookService.GetAllBooks(c)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, "Books fetched successfully", books)
}

func (r *BookRoutes) CreateBook(c echo.Context) error {
	book := &models.Book{}
	if err := c.Bind(book); err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	if book.Title == "" || book.Author == "" {
		return utils.JSONResponse(c, http.StatusBadRequest, "Title and Author are required")
	}

	book, err := r.bookService.CreateBook(c, book)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, "Book created successfully", book)
}

func (r *BookRoutes) UpdateBook(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}
	book := &models.Book{}
	if err := c.Bind(book); err != nil {

		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
	}

	// Basic validation
	if book.Title == "" || book.Author == "" || book.PublishedDate == "" {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Title and Author are required")
	}
	book.ID = idInt
	book, err = r.bookService.UpdateBook(c, book)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, "Book updated successfully", book)
}

func (r *BookRoutes) DeleteBook(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	err = r.bookService.DeleteBook(c, idInt)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, "Book deleted successfully", nil)
}

func (r *BookRoutes) GetBookByID(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	if idInt == 0 {
		return utils.ErrorResponse(c, http.StatusBadRequest, "ID is required")
	}
	book, err := r.bookService.GetBookByID(c, idInt)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, "Book fetched successfully", book)
}

func (r *BookRoutes) GetRoutes(e *echo.Echo) {
	group := e.Group("/book")

	group.GET("", r.GetAllBooks)
	group.POST("", r.CreateBook)
	group.GET("/:id", r.GetBookByID)
	group.PUT("/:id", r.UpdateBook)
	group.DELETE("/:id", r.DeleteBook)
}
