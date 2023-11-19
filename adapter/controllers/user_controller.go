package controllers

import (
	"net/http"

	services "mythos-auth/application/services/register_user"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	services.RegisterUserService
}

func NewUserController(rus services.RegisterUserService) UserController {
	return UserController{rus}
}

func (uc *UserController) Create(c echo.Context) error {
	uf := new(UserForm)

	if err := c.Bind(uf); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	od, err := uc.RegisterUserService.Invoke(&services.RegisterUserInputData{
		Email:    uf.Email,
		Password: uf.Password,
	})
	if err != nil {
		return c.String(http.StatusInternalServerError, "Internal Server Error")
	}

	return c.JSON(
		http.StatusCreated,
		newUserResponse(
			od.Id,
			od.Email,
			od.CreatedAt,
			od.UpdatedAt,
		),
	)
}
