package handler

import (
	"awesomeProject/model"
	"awesomeProject/pkg/pub_err"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 订单列表获取
func (h *Handler) GetOrders(c *gin.Context) {
	req := &PageReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		ContextResp(c, 400, &model.Result{Success: false, ErrMsg: "错误！"})
		return
	}
	filter := &model.QueryFilter{OrderBy: "created_at DESC", Page: req.Page, PageSize: req.PageSize}

	//if err := c.ShouldBindJSON(req); err != nil {
	//	ContextResp(c,400,&model.Result{
	//		Success: ,
	//		Data:
	//	})
	//}
	admins, err := h.OrderService.GetOrder(filter)
	if err != nil {
		ContextResp(c, 400, &model.Result{Success: false, ErrMsg: "错误！"})
	}
	ContextResp(c, 200, &model.Result{Success: true, Data: &model.PageResult{List: admins, TotalCount: filter.TotalCount, TotalPages: filter.TotalPage}})
}

//新增订单列表
func (h *Handler) AddOrder(c *gin.Context) {
	req := &OrderReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		ContextResp(c, http.StatusBadRequest, &model.Result{Success: false, ErrMsg: pub_err.ErrLackParams.Error()})
		return
	}
	if error := h.OrderService.AddOrder(&model.DemoModel{
		OrderNo:  req.OrderNo,
		UserName: req.UserName,
		Amount:   req.Amount,
		Status:   req.Status,
		FileUrl:  req.FileUrl,
	}); error != nil {
		ContextResp(c, http.StatusInternalServerError, &model.Result{Success: false, ErrMsg: error.Error()})
		return
	}
	ContextResp(c, http.StatusOK, &model.Result{Success: true})
	return
}

//更新订单列表
func (h *Handler) UpdateOrder(c *gin.Context) {
	req := &UpdateOrderReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		ContextResp(c, http.StatusBadRequest, &model.Result{Success: false, ErrMsg: pub_err.ErrLackParams.Error()})
		return
	}
	if error := h.OrderService.UpdateOrder(&model.DemoModel{
		OrderNo: req.OrderNo,
		FileUrl: req.FileUrl,
		Status:  req.Status,
		Amount:  req.Amount,
	}); error != nil {
		ContextResp(c, http.StatusInternalServerError, &model.Result{Success: false, ErrMsg: error.Error()})
		return
	}

	ContextResp(c, http.StatusOK, &model.Result{Success: true})
	return
}

func (h *Handler) DeleteOrder(c *gin.Context) {
	req := &DeleteOrderReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		ContextResp(c, http.StatusBadRequest, &model.Result{Success: false, ErrMsg: pub_err.ErrLackParams.Error()})
		return
	}
	if error := h.OrderService.DeleteOrder(req.Id);error != nil {
		ContextResp(c, http.StatusInternalServerError, &model.Result{Success: false, ErrMsg: error.Error()})
		return
	}
	ContextResp(c,http.StatusOK,&model.Result{Success: true})
	return
}
