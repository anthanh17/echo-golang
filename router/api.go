package router

import (
	"goEcho/handler"

	"github.com/labstack/echo/v4"
)

type Api struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *Api) SetupRouter() {
	api.Echo.GET("/user/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.GET("/user/sign-up", api.UserHandler.HandleSignUp)
}
