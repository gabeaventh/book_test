package services

import (
	"book_test/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nedpals/supabase-go"
)

type UserService interface {
	SignUp(email, password string, c echo.Context) (*supabase.AuthenticatedDetails, error)
	SignIn(email, password string, c echo.Context) (*supabase.AuthenticatedDetails, error)
	SignOut(c echo.Context) error
	GetUser(token string, c echo.Context) (*supabase.User, error)
}

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &UserServiceImpl{userRepository: userRepository}
}

func (u *UserServiceImpl) SignUp(email, password string, c echo.Context) (*supabase.AuthenticatedDetails, error) {
	if email == "" || password == "" {
		return nil, echo.NewHTTPError(400, "Email and password are required")
	}
	return u.userRepository.SignUp(email, password, c)
}

func (u *UserServiceImpl) SignIn(email, password string, c echo.Context) (*supabase.AuthenticatedDetails, error) {
	if email == "" || password == "" {
		return nil, echo.NewHTTPError(400, "Email or password are wrong")
	}
	return u.userRepository.SignIn(email, password, c)
}

func (u *UserServiceImpl) SignOut(c echo.Context) error {
	if c.Request().Header.Get("Authorization") == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}
	return u.userRepository.SignOut(c)
}

func (u *UserServiceImpl) GetUser(token string, c echo.Context) (*supabase.User, error) {
	if token == "" {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}
	return u.userRepository.GetUser(token, c)
}
