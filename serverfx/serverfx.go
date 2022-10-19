package serverfx

import (
	"context"
	"gin-fx/common"
	"gin-fx/handlers"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var(
	AppServer *appServer
	one *handlers.Handler
	two *handlers.Handler1
)

type appServer struct {
	HttpServer *gin.Engine
}

func NewAppServer(httpServ *gin.Engine) *appServer{
	return &appServer{
		HttpServer: httpServ,
	}
}

func InitServer(lc fx.Lifecycle) {
	httpServ := http.Server{
		Addr:    _port,
		Handler: AppServer.HttpServer,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				c := make(chan os.Signal)
				signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
				<-c
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				err := httpServ.Shutdown(ctx)
				if err != nil {
					cancel()
				}
			}()
			common.ManagerContainer.Logger.Info("Start HttpServer "+_port)
			go httpServ.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			common.ManagerContainer.Logger.Info("OnStop Called")
			return nil
		},
	})
}

var Module = fx.Options(
	fx.Provide(
		handlers.NewHandler,
		handlers.NewHandler1,
		NewHttpServer,
		NewAppServer,
	),
	fx.Populate(&one,&two,&AppServer),
	fx.Invoke(
		InitServer,
	),
)
