package server

import (
	"fmt"
	"net/http"

	"mythos-auth/adapter/controllers"
	"mythos-auth/application/change_email"
	"mythos-auth/application/register_user"

	"mythos-auth/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("Request Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
}

func NewRouter() (*echo.Echo, error) {
	c := config.GetConfig()
	router := echo.New()
	// router.Use(middleware.BodyDump(bodyDumpHandler))
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
	mailController := controllers.NewMailController()
	loginFlowAfterController := controllers.NewLoginFlowAfterController()

	version.POST("/users", userController.Create)
	version.PUT("/users/:id/email", userController.ChangeEmail)
	version.POST("/mail/send", mailController.Send)
	version.POST("/login/flow/after", loginFlowAfterController.LoginFlowAfter)

	return router, nil
}
