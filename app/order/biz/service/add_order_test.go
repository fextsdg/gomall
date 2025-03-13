package service

import (
	"context"
	"github.com/joho/godotenv"
	"gomall/app/order/biz/dal/mysql"
	"gomall/rpc_gen/kitex_gen/cart"
	"gomall/rpc_gen/kitex_gen/checkout"
	order "gomall/rpc_gen/kitex_gen/order"
	"testing"
)

func TestAddOrder_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.InitTest1()
	ctx := context.Background()

	s := NewAddOrderService(ctx)

	// 初始化请求
	req := &order.AddOrderReq{
		UserId:       12345,
		UserCurrency: "USD",
		Email:        "test@example.com",
		Address: &checkout.Address{
			Street:  "123 Main St",
			City:    "Anytown",
			State:   "CA",
			Country: "USA",
			ZipCode: "12345",
		},
		OrderItems: []*order.OrderItem{
			&order.OrderItem{
				CartItem: &cart.CartItem{
					ProductId: 67890,
					Num:       2,
				},
				Cost: 19.99,
			},
			&order.OrderItem{
				CartItem: &cart.CartItem{
					ProductId: 67891,
					Num:       1,
				},
				Cost: 29.99,
			},
			&order.OrderItem{
				CartItem: &cart.CartItem{
					ProductId: 67892,
					Num:       3,
				},
				Cost: 9.99,
			},
		},
	}

	// 调用服务
	resp, err := s.Run(req)

	// 验证错误
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 验证响应
	if resp == nil {
		t.Errorf("Expected a response, got nil")
	} else {
		t.Log(resp)

	}
}
