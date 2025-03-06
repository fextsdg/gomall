package service

import (
	"context"
	"github.com/joho/godotenv"
	"gomall/app/cart/biz/dal/mysql"
	"gomall/app/cart/rpc"
	cart "gomall/rpc_gen/kitex_gen/cart"
	"testing"
)

func TestAddCart_Run(t *testing.T) {
	godotenv.Load("../../.env")
	ctx := context.Background()
	mysql.InitTest()
	// 初始化RPC客户端
	rpc.InitTest()

	s := NewAddCartService(ctx)

	req := &cart.AddCartReq{
		UserId: 1,
		Item:   &cart.CartItem{ProductId: 2, Num: 2},
	}
	_, err := s.Run(req)
	if err != nil {
		t.Errorf("AddCart failed: %v", err)
	}
}
