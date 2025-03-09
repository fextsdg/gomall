package service

import (
	"context"
	checkout "gomall/rpc_gen/kitex_gen/checkout"
	"testing"
)

func TestCheckOut_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCheckOutService(ctx)
	// init req and assert value

	req := &checkout.CheckOutReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
