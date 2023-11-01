package handler

import (
	"goEcho/model"
	"goEcho/model/req"
	"goEcho/repository"
	"goEcho/security"
	"net/http"

	"goEcho/log"

	validator "github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserRepo repository.UserRepo
}

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
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	// Validate
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	// Generate hash password
	hash := security.HashAndSalt([]byte(req.Password))
	role := model.MEMBER.String()

	// User ID
	userId, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	// Insert data to database
	user := model.User{
		UserId:   userId.String(),
		FullName: req.FullName,
		Email:    req.Email,
		Password: hash,
		Role:     role,
		Token:    "",
	}
	user, err = u.UserRepo.SaveUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	// If not error return token for user
	user.Password = "" // Skip password print when call api
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Xử lý thành công",
		Data:       user,
	})
}
