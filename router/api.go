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
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)
}
