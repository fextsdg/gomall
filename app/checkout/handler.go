package main

import (
	"context"
	"gomall/app/checkout/biz/service"
	"gomall/rpc_gen/kitex_gen/checkout"
)

// CheckOutServiceImpl implements the last service interface defined in the IDL.
type CheckOutServiceImpl struct{}

// CheckOut implements the CheckOutServiceImpl interface.
func (s *CheckOutServiceImpl) CheckOut(ctx context.Context, req *checkout.CheckOutReq) (resp *checkout.CheckOutResp, err error) {
	resp, err = service.NewCheckOutService(ctx).Run(req)

	return resp, err
}
