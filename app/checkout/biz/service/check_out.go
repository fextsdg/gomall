package service

import (
	"context"
	checkout "gomall/rpc_gen/kitex_gen/checkout"
)

type CheckOutService struct {
	ctx context.Context
} // NewCheckOutService new CheckOutService
func NewCheckOutService(ctx context.Context) *CheckOutService {
	return &CheckOutService{ctx: ctx}
}

// Run create note info
func (s *CheckOutService) Run(req *checkout.CheckOutReq) (resp *checkout.CheckOutResp, err error) {
	// Finish your business logic.

	return
}
