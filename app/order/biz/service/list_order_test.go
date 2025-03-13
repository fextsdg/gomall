package service

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"gomall/app/order/biz/dal/mysql"
	order "gomall/rpc_gen/kitex_gen/order"
	"testing"
	"time"
)

func TestListOrder_Run(t *testing.T) {
	_ = godotenv.Load("../../.env")
	mysql.InitTest1()
	ctx := context.Background()
	s := NewListOrderService(ctx)
	// init req and assert value

	req := &order.ListOrderReq{
		UserId: 12345,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	if resp == nil {
		t.Fatalf("resp为空！")
	}
	for i, o := range resp.Orders {
		fmt.Printf("订单%d:\n", i)
		fmt.Printf("UserId:%v\n", o.UserId)
		fmt.Printf("OrderId:%v\n", o.OrderId)
		fmt.Printf("Address:%v\n", o.Address)
		fmt.Printf("UserCurrency:%v\n", o.UserCurrency)
		fmt.Printf("OrderItems:%v\n", o.OrderItems)
		fmt.Printf("Email%v\n", o.Email)
		// 将 Unix 时间戳转换为 time.Time 类型
		createdAt := time.Unix(o.CreatedAt, 0)
		fmt.Printf("CreatedAt:%v\n", createdAt.Format("2006-01-02 15:04:05"))

	}

	// todo: edit your unit test

}
