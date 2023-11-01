package repository

import (
	"context"
	"goEcho/model"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
}
