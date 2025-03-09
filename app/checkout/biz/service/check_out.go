package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/hashicorp/go-uuid"
	"gomall/app/checkout/rpc"
	"gomall/rpc_gen/kitex_gen/cart"
	checkout "gomall/rpc_gen/kitex_gen/checkout"
	"gomall/rpc_gen/kitex_gen/payment"
	"gomall/rpc_gen/kitex_gen/product"
)

type CheckOutService struct {
	ctx context.Context
} // NewCheckOutService new CheckOutService
func NewCheckOutService(ctx context.Context) *CheckOutService {
	return &CheckOutService{ctx: ctx}
}

// Run create note info
func (s *CheckOutService) Run(req *checkout.CheckOutReq) (resp *checkout.CheckOutResp, err error) {
	// Finish your business logic.
	getCart, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: uint32(req.UserId)})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, err.Error())
	}
	if getCart.Cart == nil || getCart.Cart.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004002, "参数错误！")
	}
	var amount float32
	cartItem := getCart.Cart.Items
	for _, item := range cartItem {
		getProduct, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: int32(item.ProductId)})
		if err != nil || getProduct.Product == nil {
			continue
		}
		amount += getProduct.Product.Price * float32(item.Num)
	}
	orderId, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}
	if req.CreditInfo == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004003, "参数有误！")
	}
	chargeResp, err := rpc.PaymentClient.Charge(s.ctx, &payment.ChargeReq{
		UserId:     uint32(req.UserId),
		OrderId:    orderId,
		CreditInfo: req.GetCreditInfo(),
		Amount:     amount,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}

	return &checkout.CheckOutResp{
		OrderId:       orderId,
		TransactionId: chargeResp.GetTransactionId(),
	}, nil
}
