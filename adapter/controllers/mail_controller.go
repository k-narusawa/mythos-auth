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
	ve := new(VerificationEmailRequest)

	if err := c.Bind(ve); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	fmt.Fprintf(os.Stdout, "Recipient: %v\n", ve.Recipient)
	fmt.Fprintf(os.Stdout, "To: %v\n", ve.To)
	fmt.Fprintf(os.Stdout, "TemplateType: %v\n", ve.TemplateType)

	return c.JSON(
		http.StatusCreated,
		"{}",
	)
}
