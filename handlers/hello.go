package handlers

import (
	"gin-fx/common"
	"gin-fx/server"
	"github.com/gin-gonic/gin"
	"net/http"
)


func NewHandler() *Handler{
	return &Handler{}
}

type Handler struct {}

func (h *Handler)Default(c *gin.Context) {
	server.ServerContainer.User.UserRepo.PrintInfo()
	common.ManagerContainer.Logger.Info("信息测试 !")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong ! success",
	})
}

