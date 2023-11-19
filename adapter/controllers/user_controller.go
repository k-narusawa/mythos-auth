package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type UserController struct{}

func NewUserController() *UserController {
	return new(UserController)
}

func (uc *UserController) Create(c echo.Context) error {
	uf := new(UserForm)

	if err := c.Bind(uf); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(
		http.StatusCreated,
		newUserResponse(
			1,
			"hoge@example.com",
			time.Now(),
			time.Now(),
		),
	)
}
