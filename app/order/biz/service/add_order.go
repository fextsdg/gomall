package service

import (
	"context"
	order "gomall/rpc_gen/kitex_gen/order"
)

type AddOrderService struct {
	ctx context.Context
} // NewAddOrderService new AddOrderService
func NewAddOrderService(ctx context.Context) *AddOrderService {
	return &AddOrderService{ctx: ctx}
}

// Run create note info
func (s *AddOrderService) Run(req *order.AddOrderReq) (resp *order.AddOrderResp, err error) {
	// Finish your business logic.

	return
}
