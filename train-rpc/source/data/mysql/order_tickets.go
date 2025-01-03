package mysql

import (
	"context"
	"errors"
	"github.com/yongsuha/goTrain/train-rpc/pkg/db"
	"github.com/yongsuha/goTrain/train-rpc/source/biz"
	"gorm.io/gorm"
)

type OrderTicketConfModel struct{}

func NewOrderTicketConfModel() biz.OrderTicketConfRepo {
	return &OrderTicketConfModel{}
}

var (
	otDB *gorm.DB
)

func (ot *OrderTicketConfModel) CreateOrderTicket() error {
	otDB = db.GetDB()
	//otDB.Create(biz.OrderTicketConf{})
	if otDB == nil {
		return errors.New("the order_tickets table init failed")
	}
	return nil
}

func (ot *OrderTicketConfModel) AddOrderTicket(ctx context.Context, orderticket *biz.OrderTicketConf) error {
	result := otDB.WithContext(ctx).Create(orderticket)
	if result.Error != nil {
		return errors.New("add to tickets failed")
	}
	return nil
}

func (ot *OrderTicketConfModel) DelOrderTicket(ctx context.Context, param *biz.OTParam) error {
	result := otDB.WithContext(ctx).Where("order_id =? AND ticket_id =?", param.OrderId, param.TicketId).Delete(&biz.OrderTicketConf{})
	if result.Error != nil {
		return errors.New("failed to delete order ticket record")
	}
	if result.RowsAffected == 0 {
		return errors.New("no order ticket record found for the given OrderId and TicketId")
	}

	return nil
}

func (ot *OrderTicketConfModel) GetOrderTicketDetail(ctx context.Context, orderId int64) (*biz.OrderTicketConf, error) {
	var orderTicketDetail biz.OrderTicketConf
	result := otDB.WithContext(ctx).Where("order_id =?", orderId).First(&orderTicketDetail)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("order ticket record not found")
		}
		return nil, errors.New("failed to get order ticket detail")
	}
	return &orderTicketDetail, nil
}
