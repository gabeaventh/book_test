package repositories

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/nedpals/supabase-go"
)

type UserRepository interface {
	SignUp(email, password string, c echo.Context) (*supabase.AuthenticatedDetails, error)
	SignIn(email, password string, c echo.Context) (*supabase.AuthenticatedDetails, error)
	SignOut(c echo.Context) error
	GetUser(token string, c echo.Context) (*supabase.User, error)
}

type UserRepositoryImpl struct {
	db *supabase.Client
}

func NewUserRepository(db *supabase.Client) UserRepository {
	return &UserRepositoryImpl{db: db}
}

// SignIn implements UserRepository.
func (u *UserRepositoryImpl) SignIn(email, password string, c echo.Context) (*supabase.AuthenticatedDetails, error) {
	ctx := context.Background()
	user, err := u.db.Auth.SignIn(ctx, supabase.UserCredentials{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return nil, err
	}
	return user, nil
}

// SignUp implements UserRepository.
func (u *UserRepositoryImpl) SignUp(email, password string, c echo.Context) (*supabase.AuthenticatedDetails, error) {
	ctx := context.Background()
	_, err := u.db.Auth.SignUp(ctx, supabase.UserCredentials{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	signIn, err := u.SignIn(email, password, c)
	if err != nil {
		return nil, err
	}
	return signIn, nil
}

// SignOut implements UserRepository.
func (u *UserRepositoryImpl) SignOut(c echo.Context) error {
	userToken := c.Request().Header.Get("Authorization")
	return u.db.Auth.SignOut(c.Request().Context(), userToken)
}

// GetUser implements UserRepository.
func (u *UserRepositoryImpl) GetUser(token string, c echo.Context) (*supabase.User, error) {
	return u.db.Auth.User(c.Request().Context(), token)
}
