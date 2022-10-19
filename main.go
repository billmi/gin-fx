package main

import (
	"gin-fx/common"
	"gin-fx/server"
	"gin-fx/serverfx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		common.Module,
		server.Module,
		serverfx.Module,
		//fx.NopLogger,
	)
	app.Run()
}
