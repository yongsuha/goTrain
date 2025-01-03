package service

import (
	"context"
	"errors"
	orderTicketRpc "github.com/yongsuha/goTrain/train-proto/order_tickets"
	"github.com/yongsuha/goTrain/train-rpc/source/biz"
	"github.com/yongsuha/goTrain/train-rpc/source/data/mysql"
	"log"
	"time"
)

type OrderTicketService struct {
	orderTicketConfModel biz.OrderTicketConfRepo
}

func NewOrderTicketService() *OrderTicketService {
	return &OrderTicketService{
		orderTicketConfModel: mysql.NewOrderTicketConfModel(),
	}
}

func (ot *OrderTicketService) AddOrderTicket(ctx context.Context, req *orderTicketRpc.AddOrderTicketReq) error {
	tag := "OrderTicketService:AddOrderTicket"
	if req.TicketId == 0 {
		return errors.New("TicketId is empty")
	}
	if req.OrderId == 0 {
		return errors.New("OrderId is empty")
	}
	otInfo := &biz.OrderTicketConf{
		OrderId:    req.OrderId,
		TicketId:   req.TicketId,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err := ot.orderTicketConfModel.AddOrderTicket(ctx, otInfo)
	if err != nil {
		log.Println(ctx, tag, "AddOrderTicket", "AddOrderTicket err:%v", err)
		return err
	}
	return nil
}

func (ot *OrderTicketService) GetOTDetail(ctx context.Context, req *orderTicketRpc.GetOTDetailReq) (*orderTicketRpc.GetOTDetailResp, error) {
	tag := "OrderTicketService:GetOTDetail"
	if req.OrderId == 0 {
		return nil, errors.New("OrderId is empty")
	}
	orderTicket, err := ot.orderTicketConfModel.GetOrderTicketDetail(ctx, req.OrderId)
	if err != nil {
		log.Println(ctx, tag, "GetOTDetail", "GetOrderTicketDetail err:%v", err)
		return nil, err
	}
	if orderTicket == nil {
		return nil, errors.New("orderTickets is empty")
	}
	resp := &orderTicketRpc.GetOTDetailResp{
		OrderId:    orderTicket.OrderId,
		TicketId:   orderTicket.TicketId,
		CreateTime: orderTicket.CreateTime.Format("2006-01-02 15:04:05"),
		UpdateTime: orderTicket.UpdateTime.Format("2006-01-02 15:04:05"),
	}
	return resp, nil
}

func (ot *OrderTicketService) DelOrderTicket(ctx context.Context, req *orderTicketRpc.DelOrderTicketReq) error {
	tag := "OrderTicketService:DelOrderTicket"
	if req.TicketId == 0 {
		return errors.New("TicketId is empty")
	}
	if req.OrderId == 0 {
		return errors.New("OrderId is empty")
	}
	orderTickets, err := ot.orderTicketConfModel.GetOrderTicketDetail(ctx, req.OrderId)
	if err != nil {
		log.Println(ctx, tag, "DelOrderTicket", "GetOrderTicketDetail err:%v", err)
		return err
	}
	if orderTickets == nil {
		return errors.New("orderTickets is empty")
	}
	delOTParam := &biz.OTParam{
		OrderId:  req.OrderId,
		TicketId: req.TicketId,
	}
	err = ot.orderTicketConfModel.DelOrderTicket(ctx, delOTParam)
	if err != nil {
		log.Println(ctx, tag, "DelOrderTicket", "DelOrderTicket err:%v", err)
		return err
	}
	return nil
}
