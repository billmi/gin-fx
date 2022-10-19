package handlers

import (
	"gin-fx/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewHandler1() *Handler1{
	return &Handler1{}
}

type Handler1 struct {}

func (h *Handler1)Default(c *gin.Context) {
	common.ManagerContainer.Logger.Info("测试初始化注入成功 !")
	c.JSON(http.StatusOK, gin.H{
		"message": "test success !",
	})
}

