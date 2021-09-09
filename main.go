package main

import (
	"awesomeProject/dao"
	"awesomeProject/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//demo_order
type DemoOrder struct {
	gorm.Model
	OrderNo  string
	UserName string
	Amount   float64
	Status   string
	FileUrl  string
}

func main() {
	err := dao.InitDb()
	if err != nil {
		panic("数据库连接失败！")
	}
	router.InitRouter()
}
