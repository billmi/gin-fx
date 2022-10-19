package common

import (
	"gin-fx/common/loggerfx"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var ManagerContainer *managerContainer

type managerContainer struct {
	Logger *zap.Logger
}

func NewManagerPlug() *managerContainer{
	return &managerContainer{
		Logger: loggerfx.NewLogger(),
	}
}

var Module = fx.Options(
	fx.Provide(
		NewManagerPlug,
	),
	fx.Populate(&ManagerContainer),
)




