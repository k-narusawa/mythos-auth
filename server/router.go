package server

import (
	"net/http"

	"mythos-auth/adapter/controllers"
	services "mythos-auth/application/services/register_user"
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
	registerUserService := services.NewRegisterUserService()
	userController := controllers.NewUserController(*registerUserService)
	version.GET("/health", healthController.Index)
	version.POST("/users", userController.Create)

	return router, nil
}
