package router

import (
	"goEcho/handler"
	"goEcho/middleware"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
	RepoHandler handler.RepoHandler
}

func (api *API) SetupRouter() {
	// user
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)

	// profile
	api.Echo.GET("/profile", api.UserHandler.Profile, middleware.JWTMiddleware())
	api.Echo.PUT("/profile/update", api.UserHandler.UpdateProfile, middleware.JWTMiddleware())

	// github repo
	github := api.Echo.Group("/github", middleware.JWTMiddleware())
	github.GET("/trending", api.RepoHandler.RepoTrending)

	// bookmark
	bookmark := api.Echo.Group("/bookmark", middleware.JWTMiddleware())
	bookmark.GET("/list", api.RepoHandler.SelectBookmarks)
	bookmark.POST("/add", api.RepoHandler.Bookmark)
	bookmark.DELETE("/delete", api.RepoHandler.DelBookmark)
}
