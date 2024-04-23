package routes

import (
	"book_test/models"
	"book_test/services"
	"book_test/utils"

	"github.com/labstack/echo/v4"
)

type UserRoutes struct {
	userService services.UserService
}

func NewUserRoutes(userService services.UserService) *UserRoutes {
	return &UserRoutes{
		userService: userService,
	}
}

func (r *UserRoutes) SignUp(c echo.Context) error {
	user := &models.UserAuth{}

	if err := c.Bind(user); err != nil {
		return err
	}

	newUser, err := r.userService.SignUp(user.Email, user.Password, c)
	if err != nil {
		return err
	}
	return utils.SuccessResponse(c, "User created successfully", newUser)
}

func (r *UserRoutes) SignIn(c echo.Context) error {
	auth := &models.UserAuth{}

	if err := c.Bind(auth); err != nil {
		return err
	}

	user, err := r.userService.SignIn(auth.Email, auth.Password, c)
	if err != nil {
		return err
	}
	return utils.SuccessResponse(c, "User signed in successfully", user)
}

func (r *UserRoutes) SignOut(c echo.Context) error {
	err := r.userService.SignOut(c)
	if err != nil {
		return err
	}
	return utils.SuccessResponse(c, "User signed out successfully", nil)
}

func (r *UserRoutes) GetRoutes(e *echo.Echo) {
	group := e.Group("/user")

	group.POST("/signup", r.SignUp)
	group.POST("/signin", r.SignIn)
	group.POST("/signout", r.SignOut)
}
