package dao

import (
	"awesomeProject/model"
	"github.com/jinzhu/gorm"
)

type OrderDao interface {
	GetOrder(filter *model.QueryFilter) ([]*model.DemoModel, error)
	AddOrder(record *model.DemoModel) error
	UpdateOrder(order *model.DemoModel) error
	DeleteOrder(id uint) error
}

type OrderDaoImpl struct {
	Db *gorm.DB
}

//确保单例
func NewOrderDaoImpl() *OrderDaoImpl {
	if orderDao == nil {
		orderDao = &OrderDaoImpl{Db: DB}
	}
	return orderDao
}

//获取order
func (d *OrderDaoImpl) GetOrder(filter *model.QueryFilter) ([]*model.DemoModel, error) {
	var orders []*model.DemoModel
	sql := d.Db.Model(&model.DemoModel{})
	if filter.OrderBy != "" {
		sql = sql.Order(filter.OrderBy)
	}
	if filter.Page > 0 && filter.PageSize > 0 {
		sql.Count(&filter.TotalCount)
		filter.TotalPage = (filter.TotalCount + filter.PageSize - 1) / filter.PageSize
		sql.Offset((filter.Page - 1) * filter.PageSize).Limit(filter.PageSize)
	}
	err := sql.Find(&orders).Error
	return orders, err
}

func (d *OrderDaoImpl) AddOrder(record *model.DemoModel) error {
	return d.Db.Model(&model.DemoModel{}).Create(record).Error
}

func (d *OrderDaoImpl) UpdateOrder(order *model.DemoModel) error {
	return d.Db.Model(&model.DemoModel{}).Where("id = ?", order.OrderNo).Update(order).Error
}

func (d *OrderDaoImpl) DeleteOrder(id uint) error {
	return d.Db.Where("id = ?", id).Delete(&model.DemoModel{}).Error
}
