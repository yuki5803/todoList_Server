package handler

import (
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

//handler结构体
type Handler struct {
	OrderService service.OrderService
}

//handler构造函数
func NewHandler() *Handler {
	return &Handler{
		OrderService: service.NewOrderServiceImpl(),
	}
}

func ContextResp(context *gin.Context, status int, result interface{}) {
	context.Header("code", strconv.Itoa(status))
	context.JSON(status, result)
}
