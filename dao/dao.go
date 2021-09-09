package dao

import "github.com/jinzhu/gorm"

const dbConfig = "root:123456@tcp(127.0.0.1:3306)/skyleForum?charset=utf8&parseTime=true&loc=Local"

var (
	DB       *gorm.DB
	orderDao *OrderDaoImpl
)

//初始化数据库连接
func InitDb() error {
	db, err := gorm.Open("mysql", dbConfig)
	if err != nil {
		panic("mysql connection has failed:" + err.Error())
	}
	DB = db
	return nil
}
