package controllers

import (
	"net/http"

	"mythos-auth/application/change_email"
	"mythos-auth/application/register_user"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	RegisterUserService register_user.RegisterUserService
	ChangeEmailService  change_email.ChangeEmailService
}

func NewUserController(
	rus register_user.RegisterUserService,
	ces change_email.ChangeEmailService,
) UserController {
	return UserController{rus, ces}
}

func (uc *UserController) Create(c echo.Context) error {
	uf := new(UserForm)

	if err := c.Bind(uf); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	od, err := uc.RegisterUserService.Invoke(&register_user.RegisterUserInputData{
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

func (uc *UserController) ChangeEmail(c echo.Context) error {
	id := c.Param("id")
	cef := new(ChangeEmailForm)

	if err := c.Bind(cef); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	od, err := uc.ChangeEmailService.Invoke(&change_email.ChangeEmailInputData{
		Id:    id,
		Email: cef.Email,
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
