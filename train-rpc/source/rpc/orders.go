package rpc

import (
	"context"
	ordersRpc "github.com/yongsuha/goTrain/train-proto/orders"
	"github.com/yongsuha/goTrain/train-rpc/source/service"
)

type OrdersServer struct {
	ordersCen *service.OrdersService
}

func NewOrdersService() *OrdersServer {
	return &OrdersServer{
		ordersCen: service.NewOrdersService(),
	}
}

func (o *OrdersServer) AddOrder(ctx context.Context, req *ordersRpc.AddOrderReq) (*ordersRpc.AddOrderResp, error) {
	return o.ordersCen.AddOrder(ctx, req)
}

func (o *OrdersServer) DelOrder(ctx context.Context, req *ordersRpc.DelOrderReq) error {
	return o.ordersCen.DelOrder(ctx, req)
}

func (o *OrdersServer) GetOrderDetail(ctx context.Context, req *ordersRpc.GetOrderDetailReq) (*ordersRpc.GetOrderDetailResp, error) {
	return o.ordersCen.GetOrderDetail(ctx, req)
}

func (o *OrdersServer) UpdateOrder(ctx context.Context, req *ordersRpc.UpdateOrderReq) (*ordersRpc.UpdateOrderResp, error) {
	return o.ordersCen.UpdateOrder(ctx, req)
}
