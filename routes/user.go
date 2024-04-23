package routes

import (
	"book_test/models"
	"book_test/services"
	"book_test/utils"
	"net/http"

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
		return utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	newUser, err := r.userService.SignUp(user.Email, user.Password, c)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, "User created successfully", newUser)
}

func (r *UserRoutes) SignIn(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := r.userService.SignIn(email, password, c)
	if err != nil {
		return utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	return utils.SuccessResponse(c, "User signed in successfully", user)
}

func (r *UserRoutes) SignOut(c echo.Context) error {
	return utils.SuccessResponse(c, "User signed out successfully", nil)
}

func (r *UserRoutes) GetRoutes(e *echo.Echo) {
	group := e.Group("/user")

	group.POST("/signup", r.SignUp)
	group.POST("/signin", r.SignIn)
	group.POST("/signout", r.SignOut)
}
