package service

import (
	"context"
	cart "gomall/rpc_gen/kitex_gen/cart"
	"testing"
)

func TestAddCart_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddCartService(ctx)
	// init req and assert value

	req := &cart.AddCartReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
