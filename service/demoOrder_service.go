package service

import (
	"awesomeProject/dao"
	"awesomeProject/model"
	"errors"
)

var (
	errorNotExist    = errors.New("不存在该ID的记录")
	errorLackRequird = errors.New("参数缺失或非法")
)

type OrderService interface {
	GetOrder(filter *model.QueryFilter) ([]*model.DemoModel, error)
	AddOrder(record *model.DemoModel) error
	UpdateOrder(order *model.DemoModel) error
	DeleteOrder(id uint) error
}

//service用例
type OrderServiceImpl struct {
	OrderDao dao.OrderDao
}

func NewOrderServiceImpl() *OrderServiceImpl {
	return &OrderServiceImpl{
		OrderDao: dao.NewOrderDaoImpl(),
	}
}

func (s *OrderServiceImpl) GetOrder(filter *model.QueryFilter) ([]*model.DemoModel, error) {
	orders, err := s.OrderDao.GetOrder(filter)
	if err != nil {
		return nil, err
	}
	return orders, err
}

func (s *OrderServiceImpl) AddOrder(record *model.DemoModel) error {
	return s.OrderDao.AddOrder(record)
}

func (s *OrderServiceImpl) UpdateOrder(order *model.DemoModel) error {
	return s.OrderDao.UpdateOrder(order)
}

func (s *OrderServiceImpl) DeleteOrder(id uint) error {
	return s.OrderDao.DeleteOrder(id)
}
