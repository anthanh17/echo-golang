package handler

import (
	"goEcho/model"
	"goEcho/model/req"
	"goEcho/security"
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (u *UserHandler) HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "An Thanh",
		"email": "@gmail.com",
	})
}

func (u *UserHandler) HandleSignUp(c echo.Context) error {

	req := req.ReqSignUp{}

	// Check request
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	// Validate
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	// Generate hash password
	hash := security.HashAndSalt([]byte(req.Password))

	type User struct {
		Name  string
		Email string
	}

	user := User{
		Name:  "Nguyen The An",
		Email: "anguyenthe.work@gmail.com",
	}
	return c.JSON(http.StatusOK, user)
}
