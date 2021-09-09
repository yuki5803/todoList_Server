package main

import (
	"awesomeProject/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const dbConfig = "root:123456@tcp(127.0.0.1:3306)/skyleForum?charset=utf8&parseTime=true&loc=Local"

func main() {
	db, err := gorm.Open("mysql", dbConfig)
	if err != nil {
		panic("mysql connection has failed:" + err.Error())
	}
	db.AutoMigrate(model.DemoModel{})
	defer db.Close()

}
