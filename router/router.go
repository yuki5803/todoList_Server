package router

import (
	"awesomeProject/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	OrderHandle := handler.NewHandler()
	r := gin.Default()
	r.POST("/api/get_order", OrderHandle.GetOrders)
	r.POST("/api/add_order", OrderHandle.AddOrder)
	r.POST("/api/update_order", OrderHandle.UpdateOrder)
	r.POST("/api/delete_order",OrderHandle.DeleteOrder)
	r.Run("0.0.0.0:9999") // 监听并在 0.0.0.0:9999 上启动服务
}
