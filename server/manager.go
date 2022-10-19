package server

import (
	"gin-fx/server/user"
	"go.uber.org/fx"
)

var ServerContainer *serverContainer

type serverContainer struct {
	User *user.UserService
}

func NewServerManager(userServ *user.UserService) *serverContainer{
	return &serverContainer{
		User: userServ,
	}
}

var Module = fx.Options(
	fx.Provide(
		user.GetUserService,
		NewServerManager,
	),
	fx.Populate(&ServerContainer),
)
