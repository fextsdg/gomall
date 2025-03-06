package service

import (
	"context"
	cart "gomall/rpc_gen/kitex_gen/cart"
)

type AddCartService struct {
	ctx context.Context
} // NewAddCartService new AddCartService
func NewAddCartService(ctx context.Context) *AddCartService {
	return &AddCartService{ctx: ctx}
}

// Run create note info
func (s *AddCartService) Run(req *cart.AddCartReq) (resp *cart.AddCartResp, err error) {
	// Finish your business logic.

	return
}
