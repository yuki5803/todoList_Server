package model

//分页封装
type QueryFilter struct {
	OrderBy                               string
	TotalCount, TotalPage, PageSize, Page int
}

//分页响应
type PageResult struct {
	List       interface{} `json:"list"`
	TotalPages int         `json:"total_pages,omitempty"`
	TotalCount int         `json:"total_count,omitempty"`
}

//DemoModel实体类
type DemoModel struct {
	BaseModel
	OrderNo  string  `json:"order_no"`
	UserName string  `json:"user_name"`
	Amount   float64 `json:"amount"`
	Status   string  `json:"status"`
	FileUrl  string  `json:"file_url"`
}
