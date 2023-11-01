package repository

import (
	"context"
	"goEcho/model"
	"goEcho/model/req"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	CheckLogin(context context.Context, login req.ReqSignIn) (model.User, error)
}
