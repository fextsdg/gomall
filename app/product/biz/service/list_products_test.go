package service

import (
	"context"
	product "gomall/rpc_gen/kitex_gen/product"
	"testing"
)

func TestListProducts_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListProductsService(ctx)
	// init req and assert value

	req := &product.ListProductsReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
