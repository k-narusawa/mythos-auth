package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// HealthController controller for health request
type HealthController struct{}

// NewHealthController is constructer for HealthController
func NewHealthController() *HealthController {
	return new(HealthController)
}

// Index is index route for health
func (hc *HealthController) Index(c echo.Context) error {
	return c.JSON(http.StatusOK,
		healthResponse(
			http.StatusOK,
			http.StatusText(http.StatusOK),
			"OK",
		))
}

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func healthResponse(status int, message string, result interface{}) *response {
	return &response{status, message, result}
}
