package middleware

import (
	"goEcho/model"
	"goEcho/security"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc  {
	config := middleware.JWTConfig{
		Claims: &model.JwtCustomClaims{},
		SigningKey: security.SECRET_KEY,
	}

	return middleware.JWTConfig(config)
}