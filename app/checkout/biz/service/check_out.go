package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/nats-io/nats.go"
	"gomall/app/checkout/mq"
	"gomall/app/checkout/rpc"
	"gomall/rpc_gen/kitex_gen/cart"
	checkout "gomall/rpc_gen/kitex_gen/checkout"
	"gomall/rpc_gen/kitex_gen/email"
	"gomall/rpc_gen/kitex_gen/order"
	"gomall/rpc_gen/kitex_gen/payment"
	"gomall/rpc_gen/kitex_gen/product"
	"google.golang.org/protobuf/proto"
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
	//从购物车获取商品
	getCart, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: uint32(req.UserId)})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, err.Error())
	}
	if getCart.Cart == nil || getCart.Cart.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004002, "参数错误！")
	}
	var amount float32
	var orderItems []*order.OrderItem
	cartItem := getCart.Cart.Items
	for _, item := range cartItem {
		getProduct, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: int32(item.ProductId)})
		if err != nil || getProduct.Product == nil {
			continue
		}
		orderItems = append(orderItems, &order.OrderItem{
			CartItem: item,
			Cost:     getProduct.Product.Price * float32(item.GetNum()),
		})
		amount += getProduct.Product.Price * float32(item.Num)
	}

	//添加订单
	orderResult, err := rpc.OrderClient.AddOrder(s.ctx, &order.AddOrderReq{
		UserId:       req.UserId,
		UserCurrency: "USD",
		Email:        req.GetEmail(),
		Address:      req.GetAddress(),
		OrderItems:   orderItems,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5004004, err.Error())
	}
	orderId := orderResult.GetAddOrderResult().GetOrderId()
	if req.CreditInfo == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004003, "参数有误！")
	}
	chargeResp, err1 := rpc.PaymentClient.Charge(s.ctx, &payment.ChargeReq{
		UserId:     uint32(req.UserId),
		OrderId:    orderId,
		CreditInfo: req.GetCreditInfo(),
		Amount:     amount,
	})

	if err1 != nil || chargeResp == nil {
		return nil, err1
	}

	//生产者发送消息
	//序列化
	data, _ := proto.Marshal(&email.SendReq{
		From:        "fextsdg@example.com",
		To:          req.GetEmail(),
		ContentType: "plain/text",
		Topic:       "You have check out an order in our website!",
		Content:     "You have check out an order in our website! You can click http:localhost:8080/order to see it!",
	})
	//制定消息
	msg := nats.Msg{
		Subject: "email",
		Data:    data,
	}
	err = mq.Nc.PublishMsg(&msg) //发布订阅
	if err != nil {
		return nil, err
	}

	return &checkout.CheckOutResp{
		OrderId:       orderId,
		TransactionId: chargeResp.GetTransactionId(),
	}, nil
}
