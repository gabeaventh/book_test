package services

import (
	"book_test/models"
	"book_test/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookService interface {
	GetAllBooks(c echo.Context) ([]*models.Book, error)
	GetBookByID(c echo.Context, id int) (*models.Book, error)
	CreateBook(c echo.Context, book *models.Book) (*models.Book, error)
	UpdateBook(c echo.Context, book *models.Book) (*models.Book, error)
	DeleteBook(c echo.Context, id int) error
}

type BookServiceImpl struct {
	bookRepository repositories.BookRepository
}

func NewBookService(bookRepository repositories.BookRepository) BookService {
	return &BookServiceImpl{bookRepository: bookRepository}
}

// CreateBook implements BookService.
func (s *BookServiceImpl) CreateBook(c echo.Context, book *models.Book) (*models.Book, error) {
	if book.Title == "" || book.Author == "" || book.PublishedDate == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest)
	}

	return s.bookRepository.CreateBook(c, book)
}

// UpdateBook implements BookService.
func (s *BookServiceImpl) UpdateBook(c echo.Context, book *models.Book) (*models.Book, error) {
	if book.ID == 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest)
	}

	if book == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest)
	}

	return s.bookRepository.UpdateBook(c, book)
}

// DeleteBook implements BookService.
func (s *BookServiceImpl) DeleteBook(c echo.Context, id int) error {
	if id == 0 {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return s.bookRepository.DeleteBook(c, id)
}

// GetBookByID implements BookService.
func (s *BookServiceImpl) GetBookByID(c echo.Context, id int) (*models.Book, error) {
	if id == 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest)
	}
	return s.bookRepository.GetBookByID(c, id)
}

// GetAllBooks implements BookService.
func (s *BookServiceImpl) GetAllBooks(ctx echo.Context) ([]*models.Book, error) {
	return s.bookRepository.GetAllBooks(ctx)
}
