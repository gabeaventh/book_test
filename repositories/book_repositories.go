package repositories

import (
	"book_test/models"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type BookRepository interface {
	GetAllBooks(c echo.Context) ([]*models.Book, error)
	GetBookByID(c echo.Context, id int) (*models.Book, error)
	CreateBook(c echo.Context, book *models.Book) (*models.Book, error)
	UpdateBook(c echo.Context, book *models.Book) (*models.Book, error)
	DeleteBook(c echo.Context, id int) error
}

type BookRepositoryImpl struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &BookRepositoryImpl{db: db}
}

// CreateBook implements BookRepository.
func (b *BookRepositoryImpl) CreateBook(c echo.Context, book *models.Book) (*models.Book, error) {
	ctx := c.Request().Context()
	book.Title = strings.TrimSpace(book.Title)
	book.Author = strings.TrimSpace(book.Author)
	book.PublishedDate = strings.TrimSpace(book.PublishedDate)

	stmt, err := b.db.PrepareContext(ctx, "INSERT INTO books (title, author, published_date) VALUES ($1, $2, $3) RETURNING id, title, author, published_date")
	if err != nil {
		return nil, err
	}

	var newBook models.Book
	err = stmt.QueryRowContext(ctx, book.Title, book.Author, book.PublishedDate).Scan(&newBook.ID, &newBook.Title, &newBook.Author, &newBook.PublishedDate)
	if err != nil {
		return nil, err
	}

	return &newBook, nil
}

// DeleteBook implements BookRepository.
func (b *BookRepositoryImpl) DeleteBook(c echo.Context, id int) error {
	if id == 0 {
		return errors.New("ID is required")
	}
	ctx := c.Request().Context()
	book := models.Book{ID: id}
	err := b.db.QueryRowContext(ctx, "SELECT deleted_at FROM books WHERE id = $1", book.ID).Scan(&book.DeletedAt)
	if err != nil {
		return err
	}
	if book.DeletedAt != nil {
		return errors.New("book is not found")
	}

	stmt, err := b.db.PrepareContext(ctx, "UPDATE books SET deleted_at = $1 WHERE id = $2")
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, time.Now(), id)
	return err
}

// GetAllBooks implements BookRepository.
func (b *BookRepositoryImpl) GetAllBooks(c echo.Context) ([]*models.Book, error) {
	ctx := c.Request().Context()

	stmt, err := b.db.PrepareContext(ctx, "SELECT id, title, author, published_date, updated_at, created_at FROM books WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedDate, &book.UpdatedAt, &book.CreatedAt); err != nil {
			return nil, err
		}
		books = append(books, &book)
	}
	return books, nil
}

// GetBookByID implements BookRepository.
func (b *BookRepositoryImpl) GetBookByID(c echo.Context, id int) (*models.Book, error) {
	if id == 0 {
		return nil, errors.New("ID is required")
	}

	ctx := c.Request().Context()
	stmt, err := b.db.PrepareContext(ctx, "SELECT id, title, author, published_date, updated_at, created_at FROM books WHERE id = $1 AND deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var book models.Book
	err = stmt.QueryRowContext(ctx, id).Scan(&book.ID, &book.Title, &book.Author, &book.PublishedDate, &book.UpdatedAt, &book.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// UpdateBook implements BookRepository.
func (b *BookRepositoryImpl) UpdateBook(c echo.Context, book *models.Book) (*models.Book, error) {
	if book == nil {
		return nil, errors.New("book data cannot be empty")
	}

	ctx := c.Request().Context()
	book.Title = strings.TrimSpace(book.Title)
	book.Author = strings.TrimSpace(book.Author)
	book.PublishedDate = strings.TrimSpace(book.PublishedDate)
	err := b.db.QueryRowContext(ctx, "SELECT deleted_at FROM books WHERE id = $1", book.ID).Scan(&book.DeletedAt)
	if err != nil {
		return nil, err
	}
	if book.DeletedAt != nil {
		return nil, errors.New("book is not found")
	}
	stmt, err := b.db.PrepareContext(ctx, "UPDATE books SET title = $1, author = $2, published_date = $3, updated_at = $4 WHERE id = $5")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, book.Title, book.Author, book.PublishedDate, time.Now(), book.ID)
	if err != nil {
		return nil, err
	}
	return book, nil
}
