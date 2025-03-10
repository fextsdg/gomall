package service

import (
	"context"
	"gomall/app/checkout/rpc"
	checkout "gomall/rpc_gen/kitex_gen/checkout"
	"gomall/rpc_gen/kitex_gen/payment"
	"testing"
)

func TestCheckOut_Run(t *testing.T) {
	rpc.InitTest1()
	ctx := context.Background()
	s := NewCheckOutService(ctx)
	// init req and assert value

	req := &checkout.CheckOutReq{
		UserId:    8,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Address: &checkout.Address{
			Street:  "123 Main St",
			City:    "Anytown",
			State:   "CA",
			Country: "USA",
			ZipCode: "12345",
		},
		CreditInfo: &payment.CreditInfo{
			CreditCardNumber:     "4111111111111111",
			CreditCardCvv:        13,
			CreditExpirationYear: 2025,
			CreditExpirationMoth: 12,
		},
	}

	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	if resp == nil {
		t.Fatalf("resp为nil！")
	}
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
