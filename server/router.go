package server

import (
	"net/http"

	"mythos-auth/adapter/controllers"
	"mythos-auth/application/change_email"
	"mythos-auth/application/register_user"

	"mythos-auth/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() (*echo.Echo, error) {
	c := config.GetConfig()
	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: c.GetStringSlice("server.cors"),
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	version := router.Group("/api/" + c.GetString("server.version"))

	healthController := controllers.NewHealthController()
	version.GET("/health", healthController.Index)

	registerUserService := register_user.NewRegisterUserService()
	changeEmailService := change_email.NewChangeEmailService()
	userController := controllers.NewUserController(*registerUserService, *changeEmailService)
	version.POST("/users", userController.Create)
	version.PUT("/users/:id/email", userController.ChangeEmail)

	return router, nil
}
