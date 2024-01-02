package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginFlowAfterController struct{}

func NewLoginFlowAfterController() LoginFlowAfterController {
	return LoginFlowAfterController{}
}

func (lc *LoginFlowAfterController) LoginFlowAfter(c echo.Context) error {
	lf := new(LoginFlowAfterRequest)

	if err := c.Bind(lf); err != nil {
		fmt.Printf("Error: %v\n", err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	fmt.Printf("UserID: %v\n", lf.UserID)
	if lf.Email != nil {
		fmt.Printf("Email: %v\n", *lf.Email)
	}
	if lf.XForwardedFor != nil {
		fmt.Printf("X-Forwarded-For: %v\n", *lf.XForwardedFor)
	}

	return c.NoContent(http.StatusNoContent)
}
