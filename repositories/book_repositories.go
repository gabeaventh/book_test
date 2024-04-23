package repositories

import (
	"book_test/models"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nedpals/supabase-go"
)

const (
	defaultTime = "0001-01-01T00:00:00Z"
	layout      = "2006-01-02T15:04:05.999999+00:00"
)

type BookRepository interface {
	GetAllBooks(c echo.Context) ([]*models.Book, error)
	GetBookByID(c echo.Context, id int) (*models.Book, error)
	CreateBook(c echo.Context, book *models.Book) (*models.Book, error)
	UpdateBook(c echo.Context, book *models.Book) (*models.Book, error)
	DeleteBook(c echo.Context, id int) error
}

type BookRepositoryImpl struct {
	db *supabase.Client
}

func NewBookRepository(db *supabase.Client) BookRepository {
	return &BookRepositoryImpl{db: db}
}

// CreateBook implements BookRepository.
func (b *BookRepositoryImpl) CreateBook(c echo.Context, book *models.Book) (*models.Book, error) {
	book.Title = strings.TrimSpace(book.Title)
	book.Author = strings.TrimSpace(book.Author)
	book.PublishedDate = strings.TrimSpace(book.PublishedDate)

	// manually get the next ID
	// supabase somehow doesn't increment the ID
	var result []map[string]interface{}

	err := b.db.DB.From("books").Select("*").Execute(&result)

	if err != nil {
		return nil, err
	}

	// increment the ID
	count := len(result)
	id := 1
	if count > 0 {
		id = count + 1
	}
	newBook := &models.Book{
		ID:            id,
		Title:         book.Title,
		Author:        book.Author,
		PublishedDate: book.PublishedDate,
		CreatedAt:     time.Now(),
		UpdatedAt:     book.UpdatedAt,
		DeletedAt:     book.DeletedAt,
	}

	var results []map[string]interface{}
	err = b.db.DB.From("books").Insert(newBook).Execute(&results)

	if err != nil {
		return nil, err
	}

	return newBook, nil
}

// DeleteBook implements BookRepository.
func (b *BookRepositoryImpl) DeleteBook(c echo.Context, id int) error {
	if id == 0 {
		return errors.New("ID is required")
	}

	book, err := b.GetBookByID(c, id)
	if err != nil {
		return err
	}

	newBook := &models.Book{
		ID:            book.ID,
		Title:         book.Title,
		Author:        book.Author,
		PublishedDate: book.PublishedDate,
		CreatedAt:     book.CreatedAt,
		UpdatedAt:     book.UpdatedAt,
		DeletedAt:     time.Now(),
	}
	var res []models.Book
	err = b.db.DB.From("books").Update(newBook).Eq("id", strconv.Itoa(newBook.ID)).Execute(&res)

	if err != nil {
		return err
	}
	return nil
}

// GetAllBooks implements BookRepository.
func (b *BookRepositoryImpl) GetAllBooks(c echo.Context) ([]*models.Book, error) {
	var res []map[string]interface{}
	err := b.db.DB.From("books").Select("*").Execute(&res)

	if err != nil {
		return nil, err
	}

	var newBook []*models.Book

	for _, book := range res {
		createdAt, _ := time.Parse(layout, book["created_at"].(string))
		updatedAt, _ := time.Parse(layout, book["updated_at"].(string))
		deletedAt, _ := time.Parse(layout, book["deleted_at"].(string))

		newBook = append(newBook, &models.Book{
			ID:            int(book["id"].(float64)),
			Title:         book["title"].(string),
			Author:        book["author"].(string),
			PublishedDate: book["published_date"].(string),
			CreatedAt:     createdAt,
			UpdatedAt:     updatedAt,
			DeletedAt:     deletedAt,
		})
	}

	return newBook, nil
}

// GetBookByID implements BookRepository.
func (b *BookRepositoryImpl) GetBookByID(c echo.Context, id int) (*models.Book, error) {
	if id == 0 {
		return nil, errors.New("ID is required")
	}
	var res map[string]interface{}
	err := b.db.DB.From("books").Select("*").Single().Eq("id", strconv.Itoa(id)).Execute(&res)

	if err != nil {
		return nil, err
	}

	createAt, _ := time.Parse(layout, res["created_at"].(string))
	updatedAt, _ := time.Parse(layout, res["updated_at"].(string))
	deletedAt, _ := time.Parse(layout, res["deleted_at"].(string))
	newBook := &models.Book{
		ID:            int(res["id"].(float64)),
		Title:         res["title"].(string),
		Author:        res["author"].(string),
		PublishedDate: res["published_date"].(string),
		CreatedAt:     createAt,
		UpdatedAt:     updatedAt,
		DeletedAt:     deletedAt,
	}

	return newBook, nil
}

// UpdateBook implements BookRepository.
func (b *BookRepositoryImpl) UpdateBook(c echo.Context, book *models.Book) (*models.Book, error) {
	if book == nil {
		return nil, errors.New("book data cannot be empty")
	}
	book.Title = strings.TrimSpace(book.Title)
	book.Author = strings.TrimSpace(book.Author)
	book.PublishedDate = strings.TrimSpace(book.PublishedDate)

	newBook := &models.Book{
		ID:            book.ID,
		Title:         book.Title,
		Author:        book.Author,
		PublishedDate: book.PublishedDate,
		CreatedAt:     book.CreatedAt,
		UpdatedAt:     time.Now(),
		DeletedAt:     book.DeletedAt,
	}

	var res []*models.Book
	err := b.db.DB.From("books").Update(newBook).Eq("id", strconv.Itoa(newBook.ID)).Execute(&res)

	if err != nil {
		return nil, err
	}

	return newBook, nil
}
