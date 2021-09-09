package handler

type PageReq struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

type OrderReq struct {
	OrderNo  string  `json:"order_no"`
	UserName string  `json:"user_name"`
	Amount   float64 `json:"amount"`
	Status   string  `json:"status"`
	FileUrl  string  `json:"file_url"`
}

type UpdateOrderReq struct {
	OrderNo string  `json:"order_no" binding:"required"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"`
	FileUrl string  `json:"file_url"`
}

type DeleteOrderReq struct {
	Id uint `json:"id" binding:"required"`
}
