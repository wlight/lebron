// Code generated by goctl. DO NOT EDIT!
// Source: order.proto

package server

import (
	"context"

	"wlight/lebron/apps/order/rpc/internal/logic"
	"wlight/lebron/apps/order/rpc/internal/svc"
	"wlight/lebron/apps/order/rpc/order"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	order.UnimplementedOrderServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServer) Orders(ctx context.Context, in *order.OrdersRequest) (*order.OrdersResponse, error) {
	l := logic.NewOrdersLogic(ctx, s.svcCtx)
	return l.Orders(in)
}
