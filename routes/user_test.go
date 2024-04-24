package routes_test

import (
	"book_test/routes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/nedpals/supabase-go"
	"github.com/stretchr/testify/assert"
)

type MockUserService struct {
	mockSignUp  func(email, password string, c echo.Context) (*supabase.AuthenticatedDetails, error)
	mockSignIn  func(email, password string, c echo.Context) (*supabase.AuthenticatedDetails, error)
	mockSignOut func(c echo.Context) error
	mockGetUser func(token string, c echo.Context) (*supabase.User, error)
}

// Implement the UserService Interface
func (m *MockUserService) SignUp(email, password string, c echo.Context) (*supabase.AuthenticatedDetails, error) {
	return m.mockSignUp(email, password, c)
}

func (m *MockUserService) SignIn(email, password string, c echo.Context) (*supabase.AuthenticatedDetails, error) {
	return m.mockSignIn(email, password, c)
}

func (m *MockUserService) SignOut(c echo.Context) error {
	return m.mockSignOut(c)
}

func (m *MockUserService) GetUser(token string, c echo.Context) (*supabase.User, error) {
	return m.mockGetUser(token, c)
}

func TestUserRoutes_SignUp(t *testing.T) {
	userId := "newId"
	// Test Cases
	testCases := []struct {
		name         string
		mockBehavior func(mock *MockUserService)
		requestBody  string
		expectedCode int
		expectedBody string
	}{
		{
			name: "Success",
			mockBehavior: func(mock *MockUserService) {
				mock.mockSignUp = func(email, password string, c echo.Context) (*supabase.AuthenticatedDetails, error) {
					return &supabase.AuthenticatedDetails{
						User: supabase.User{ID: userId, Email: "test@example.com"},
					}, nil
				}
			},
			requestBody:  `{"email": "test@example.com", "password": "password123"}`,
			expectedCode: http.StatusOK,
			expectedBody: `{"status":200,"message":"User created successfully","data":{"access_token":"","token_type":"","expires_in":0,"refresh_token":"","user":{"id":"newId","aud":"","role":"","email":"test@example.com","invited_at":"0001-01-01T00:00:00Z","confirmed_at":"0001-01-01T00:00:00Z","confirmation_sent_at":"0001-01-01T00:00:00Z","app_metadata":{},"user_metadata":null,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"},"provider_token":"","provider_refresh_token":""}}`,
		},
		{
			name: "Empty Email",
			mockBehavior: func(mock *MockUserService) {
				mock.mockSignUp = func(email, password string, c echo.Context) (*supabase.AuthenticatedDetails, error) {
					return nil, fmt.Errorf("Email cannot be empty")
				}
			},
			requestBody:  `{"email": "", "password": "password123"}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: `Email cannot be empty`,
		},
		{
			name: "Supabase Error",
			mockBehavior: func(mock *MockUserService) {
				mock.mockSignUp = func(email, password string, c echo.Context) (*supabase.AuthenticatedDetails, error) {
					return nil, fmt.Errorf("Supabase signup error")
				}
			},
			requestBody:  `{"email": "test@example.com", "password": "password123"}`,
			expectedCode: http.StatusInternalServerError,
			expectedBody: `Supabase signup error`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Setup
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/user/signup", strings.NewReader(testCase.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			mockService := &MockUserService{}
			testCase.mockBehavior(mockService)
			routes := routes.NewUserRoutes(mockService)

			if err := routes.SignUp(c); err != nil {
				assert.Error(t, err)
				if httpErr, ok := err.(*echo.HTTPError); ok {
					assert.Equal(t, testCase.expectedCode, httpErr.Code)
				}
				assert.Contains(t, err.Error(), testCase.expectedBody)
			} else {
				assert.Equal(t, testCase.expectedCode, rec.Code)
				assert.Equal(t, testCase.expectedBody, strings.TrimSpace(rec.Body.String()))
			}
		})
	}
}

func TestUserRoutes_SignIn(t *testing.T) {
	userId := "newId"
	// Test Cases
	testCases := []struct {
		name         string
		mockBehavior func(mock *MockUserService)
		requestBody  string
		expectedCode int
		expectedBody string
	}{
		{
			name: "Success",
			mockBehavior: func(mock *MockUserService) {
				mock.mockSignIn = func(email, password string, c echo.Context) (*supabase.AuthenticatedDetails, error) {
					return &supabase.AuthenticatedDetails{
						User: supabase.User{ID: userId, Email: "test@example.com"},
					}, nil
				}
			},
			requestBody:  `{"email": "test@example.com", "password": "password123"}`,
			expectedCode: http.StatusOK,
			expectedBody: `{"status":200,"message":"User signed in successfully","data":{"access_token":"","token_type":"","expires_in":0,"refresh_token":"","user":{"id":"newId","aud":"","role":"","email":"test@example.com","invited_at":"0001-01-01T00:00:00Z","confirmed_at":"0001-01-01T00:00:00Z","confirmation_sent_at":"0001-01-01T00:00:00Z","app_metadata":{},"user_metadata":null,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"},"provider_token":"","provider_refresh_token":""}}`,
		},
		{
			name: "Empty Email",
			mockBehavior: func(mock *MockUserService) {
				mock.mockSignIn = func(email, password string, c echo.Context) (*supabase.AuthenticatedDetails, error) {
					return nil, fmt.Errorf("Email cannot be empty")
				}
			},
			requestBody:  `{"email": "", "password": "password123"}`,
			expectedCode: http.StatusBadRequest,
			expectedBody: `Email cannot be empty`,
		},
		{
			name: "Supabase Error",
			mockBehavior: func(mock *MockUserService) {
				mock.mockSignIn = func(email, password string, c echo.Context) (*supabase.AuthenticatedDetails, error) {
					return nil, fmt.Errorf("Supabase signup error")
				}
			},
			requestBody:  `{"email": "test@example.com", "password": "password123"}`,
			expectedCode: http.StatusInternalServerError,
			expectedBody: `Supabase signup error`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Setup
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/user/signup", strings.NewReader(testCase.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			mockService := &MockUserService{}
			testCase.mockBehavior(mockService)
			routes := routes.NewUserRoutes(mockService)

			if err := routes.SignIn(c); err != nil {
				assert.Error(t, err)
				if httpErr, ok := err.(*echo.HTTPError); ok {
					assert.Equal(t, testCase.expectedCode, httpErr.Code)
				}
				assert.Contains(t, err.Error(), testCase.expectedBody)
			} else {
				assert.Equal(t, testCase.expectedCode, rec.Code)
				assert.Equal(t, testCase.expectedBody, strings.TrimSpace(rec.Body.String()))
			}
		})
	}
}

func TestUserRoutes_SignOut(t *testing.T) {
	// Test Cases
	testCases := []struct {
		name         string
		mockBehavior func(mock *MockUserService)
		expectedCode int
		expectedBody string
	}{
		{
			name: "Success",
			mockBehavior: func(mock *MockUserService) {
				mock.mockSignOut = func(c echo.Context) error {
					return nil
				}
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"status":200,"message":"User signed out successfully"}`,
		},
		{
			name: "Supabase Error",
			mockBehavior: func(mock *MockUserService) {
				mock.mockSignOut = func(c echo.Context) error {
					return fmt.Errorf("Supabase signout error")
				}
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: `Supabase signout error`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Setup
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/user/signout", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			mockService := &MockUserService{}
			testCase.mockBehavior(mockService)
			routes := routes.NewUserRoutes(mockService)

			if err := routes.SignOut(c); err != nil {
				assert.Error(t, err)
				if httpErr, ok := err.(*echo.HTTPError); ok {
					assert.Equal(t, testCase.expectedCode, httpErr.Code)
				}
				assert.Contains(t, err.Error(), testCase.expectedBody)
			} else {
				assert.Equal(t, testCase.expectedCode, rec.Code)
				assert.Equal(t, testCase.expectedBody, strings.TrimSpace(rec.Body.String()))
			}
		})
	}
}
