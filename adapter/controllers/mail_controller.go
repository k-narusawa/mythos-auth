package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type MailController struct{}

func NewMailController() MailController {
	return MailController{}
}

func (mc *MailController) Send(c echo.Context) error {
	fmt.Fprintf(os.Stderr, "Full HTTP Request: %v\n", c.Request().Body)

	return c.JSON(
		http.StatusCreated,
		"{}",
	)
}
