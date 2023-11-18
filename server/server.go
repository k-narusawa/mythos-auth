package server

import (
	"mythos-auth/config"

	"github.com/labstack/echo/v4"
)

// Init initialize server
func Init() error {
	c := config.GetConfig()
	e := echo.New()

	e.Logger.Fatal(e.Start(":" + c.GetString("server.port")))

	return nil
}
