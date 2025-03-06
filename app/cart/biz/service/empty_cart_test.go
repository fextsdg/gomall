package service

import (
	"context"
	"github.com/joho/godotenv"
	"gomall/app/cart/biz/dal/mysql"
	"gomall/app/cart/rpc"
	cart "gomall/rpc_gen/kitex_gen/cart"
	"testing"
)

func TestEmptyCart_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.InitTest()
	rpc.InitTest()
	ctx := context.Background()
	s := NewEmptyCartService(ctx)
	// init req and assert value

	req := &cart.EmptyCartReq{UserId: 1}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
