package mysql

import (
	"context"
	"errors"
	"github.com/yongsuha/goTrain/train-rpc/source/biz"
	"github.com/yongsuha/goTrain/train-rpc/source/db"
	"gorm.io/gorm"
)

type TicketsConfModel struct{}

func NewTicketsConfModel() biz.TicketsConfRepo {
	return &TicketsConfModel{}
}

var (
	ticketDB *gorm.DB
)

func (t *TicketsConfModel) CreateTicket() error {
	ticketDB = db.GetDB()
	ticketDB.Create(biz.TicketsConf{})
	if ticketDB == nil {
		return errors.New("the tickets table init failed")
	}
	return nil
}

func (t *TicketsConfModel) AddTicket(ctx context.Context, ticket *biz.TicketsConf) (int64, error) {
	result := ticketDB.WithContext(ctx).Create(ticket)
	if result.Error != nil {
		return 0, errors.New("add to tickets failed")
	}
	return ticket.TicketId, nil
}

func (t *TicketsConfModel) DelTicket(ctx context.Context, ticketId int64) error {
	result := ticketDB.WithContext(ctx).Where("ticket_id =?", ticketId).Delete(biz.TicketsConf{})
	if result.Error != nil {
		return errors.New("delete ticket record failed")
	}
	if result.RowsAffected == 0 {
		return errors.New("no ticket record found for the given ticket_id")
	}
	return nil
}

func (t *TicketsConfModel) GetTicketDetail(ctx context.Context, param *biz.GetTicketDetailParam) ([]*biz.TicketsConf, error) {
	var tickets []*biz.TicketsConf
	query := ticketDB.WithContext(ctx)

	if param.TicketId != 0 {
		query = query.Where("ticket_id =?", param.TicketId)
	}
	if param.UserId != 0 {
		query = query.Where("user_id =?", param.UserId)
	}

	result := query.Find(&tickets)
	if result.Error != nil {
		return nil, errors.New("failed to get ticket details")
	}
	return tickets, nil
}

func (t *TicketsConfModel) UpdateTicket(ctx context.Context, ticket *biz.TicketsConf) error {
	result := ticketDB.WithContext(ctx).Model(biz.TicketsConf{}).Where("ticket_id =?", ticket.TicketId).Updates(ticket)
	if result.Error != nil {
		return errors.New("update ticket record failed")
	}
	if result.RowsAffected == 0 {
		return errors.New("no ticket record found for the given ticket_id")
	}
	return nil
}
