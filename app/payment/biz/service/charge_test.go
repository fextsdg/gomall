package service

import (
	"context"
	"github.com/joho/godotenv"
	"gomall/app/payment/biz/dal/mysql"
	payment "gomall/rpc_gen/kitex_gen/payment"
	"testing"
)

func TestCharge_Run(t *testing.T) {
	godotenv.Load("../../.env")
	ctx := context.Background()
	s := NewChargeService(ctx)
	// init req and assert value
	mysql.InitTest()
	req := &payment.ChargeReq{
		UserId:  8,              // 填充一个有效的用户ID
		OrderId: "ORD123456789", // 填充一个有效的订单ID
		CreditInfo: &payment.CreditInfo{
			CreditCardNumber:     "4111111111111111", // 填充一个有效的信用卡号
			CreditCardCvv:        123,                // 填充一个有效的CVV码
			CreditExpirationYear: 2025,               // 填充一个有效的过期年份
			CreditExpirationMoth: 12,                 // 填充一个有效的过期月份
		},
		Amount: 100.50, // 填充一个有效的金额
	}

	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// 添加断言来验证响应和错误
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp == nil {
		t.Errorf("Expected a response, got nil")
	}
	t.Logf(resp.TransactionId)

}
