package model

import (
	"time"
)

//基本表结构体
type BaseModel struct {
	ID        uint      `gorm:"column:id;index;primary_key;" json:"id,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at,omitempty"`
	DeletedAt *time.Time `gorm:"column:deleted_at;" json: "deleted_at,omitempty"`
	UpdatedAt time.Time `gorm:"column: updated_at;" json:" "updated_at,omitempty"`
}

//响应结构体
type Result struct {
	ErrMsg  string      `json:"err_msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
}
