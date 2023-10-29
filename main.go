package main

import (
	"goEcho/handler"

	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    e.GET("/", handler.Welcome)

	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)


    e.Logger.Fatal(e.Start(":3000"))
}
