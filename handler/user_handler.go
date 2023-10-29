package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user" : "An Thanh",
		"email" : "@gmail.com",
	})
}

func HandleSignUp(c echo.Context) error {
	type User struct {
		Name string
		Email string
	}

	user := User {
		Name : "Nguyen The An",
		Email: "anguyenthe.work@gmail.com",
	}
	return c.JSON(http.StatusOK, user)
}
