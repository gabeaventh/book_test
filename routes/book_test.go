package routes_test

import (
	"book_test/models"
	"book_test/routes"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
)

type MockBookService struct {
	mockGetAllBooks func(c echo.Context) ([]*models.Book, error)
	mockGetBookByID func(c echo.Context, id int) (*models.Book, error)
	mockCreateBook  func(c echo.Context, book *models.Book) (*models.Book, error)
	mockUpdateBook  func(c echo.Context, book *models.Book) (*models.Book, error)
	mockDeleteBook  func(c echo.Context, id int) error
}

func (m *MockBookService) GetAllBooks(c echo.Context) ([]*models.Book, error) {
	return m.mockGetAllBooks(c)
}

func (m *MockBookService) GetBookByID(c echo.Context, id int) (*models.Book, error) {
	return m.mockGetBookByID(c, id)
}

func (m *MockBookService) CreateBook(c echo.Context, book *models.Book) (*models.Book, error) {
	return m.mockCreateBook(c, book)
}

func (m *MockBookService) UpdateBook(c echo.Context, book *models.Book) (*models.Book, error) {
	return m.mockUpdateBook(c, book)
}

func (m *MockBookService) DeleteBook(c echo.Context, id int) error {
	return m.mockDeleteBook(c, id)
}

func NewMockBookService() *MockBookService {
	return &MockBookService{
		mockGetAllBooks: func(c echo.Context) ([]*models.Book, error) {
			return nil, nil
		},
		mockGetBookByID: func(c echo.Context, id int) (*models.Book, error) {
			return nil, nil
		},
		mockCreateBook: func(c echo.Context, book *models.Book) (*models.Book, error) {
			return nil, nil
		},
		mockUpdateBook: func(c echo.Context, book *models.Book) (*models.Book, error) {
			return nil, nil
		},
		mockDeleteBook: func(c echo.Context, id int) error {
			return nil
		},
	}
}

func TestUserRoutes_GetAllBooks(t *testing.T) {
	testCases := []struct {
		name         string
		mockBehavior func(mock *MockBookService)
		requestBody  string
		expectedCode int
		expectedBody string
	}{
		{
			name: "Success",
			mockBehavior: func(mock *MockBookService) {
				mock.mockGetAllBooks = func(c echo.Context) ([]*models.Book, error) {
					return []*models.Book{
						{ID: 1, Title: "Test Book", Author: "Test Author", PublishedDate: time.Now().Format("2006-01-02")},
					}, nil
				}
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"status":200,"message":"Books fetched successfully","data":[{"id":1,"title":"Test Book","author":"Test Author","published_date":` + time.Now().Format("2006-01-02") + `,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","deleted_at":"0001-01-01T00:00:00Z"}]}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := NewMockBookService()
			mockUser := &MockUserService{}
			tc.mockBehavior(mockService)

			// Create a new BookRoutes instance with the mock service
			routes := routes.NewBookRoutes(mockService, mockUser)

			// Create a new Echo instance
			e := echo.New()

			// Create a new HTTP request
			req := httptest.NewRequest(http.MethodGet, "/book", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Call the GetAllBooks method
			err := routes.GetAllBooks(c)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if rec.Code != tc.expectedCode {
				t.Errorf("expected status code %d, got %d", tc.expectedCode, rec.Code)
			}

			// Regular expression pattern to match your JSON with a flexible date
			jsonPattern := `{"status":200,"message":"Books fetched successfully","data":\[{"id":1,"title":"Test Book","author":"Test Author","published_date":"\d{4}-\d{2}-\d{2}","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","deleted_at":"0001-01-01T00:00:00Z"}\]}`

			match, err := regexp.MatchString(jsonPattern, rec.Body.String())
			if err != nil {
				t.Errorf("error checking JSON body: %v", err)
			}

			if !match {
				t.Errorf("expected body JSON pattern did not match. Expected pattern: %s, Got: %s", jsonPattern, rec.Body.String())
			}
		})
	}
}
