package mysql

import (
	"context"
	"errors"
	"github.com/yongsuha/goTrain/train-rpc/source/biz"
	"github.com/yongsuha/goTrain/train-rpc/source/db"
	"gorm.io/gorm"
)

type OrdersConfModel struct{}

func NewOrdersConfModel() biz.OrdersConfRepo {
	return &OrdersConfModel{}
}

var (
	orderDB *gorm.DB
)

func (o *OrdersConfModel) CreateOrder() error {
	orderDB = db.GetDB()
	orderDB.Create(biz.OrdersConf{})
	if orderDB == nil {
		return errors.New("the orders table init failed")
	}
	return nil
}

func (o *OrdersConfModel) AddOrder(ctx context.Context, order *biz.OrdersConf) (int64, error) {
	result := orderDB.WithContext(ctx).Create(order)
	if result.Error != nil {
		return 0, errors.New("add to orders failed")
	}
	return order.OrderId, nil
}

func (o *OrdersConfModel) DelOrder(ctx context.Context, orderId int64) error {
	result := orderDB.WithContext(ctx).Where("order_id =?", orderId).Delete(biz.OrdersConf{})
	if result.Error != nil {
		return errors.New("delete order record failed")
	}
	if result.RowsAffected == 0 {
		return errors.New("no order record found for the given order_id")
	}
	return nil
}

func (o *OrdersConfModel) GetOrderDetail(ctx context.Context, param *biz.GetOrderDetailParam) (*biz.OrdersConf, error) {
	order := biz.OrdersConf{}
	query := orderDB.WithContext(ctx)

	if param.OrderId != 0 {
		query = query.Where("order_id =?", param.OrderId)
	}
	if param.UserId != 0 {
		query = query.Where("user_id =?", param.UserId)
	}

	result := query.First(&order)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("order record not found")
		}
		return nil, errors.New("failed to get order detail")
	}
	return &order, nil
}

func (o *OrdersConfModel) UpdateOrder(ctx context.Context, order *biz.OrdersConf) error {
	result := orderDB.WithContext(ctx).Model(biz.OrdersConf{}).Where("order_id =?", order.OrderId).Updates(order)
	if result.Error != nil {
		return errors.New("update order record failed")
	}
	if result.RowsAffected == 0 {
		return errors.New("no order record found for the given order_id")
	}
	return nil
}
