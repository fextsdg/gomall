package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"gomall/app/order/biz/dal/mysql"
	"gomall/app/order/biz/model"
	"gomall/rpc_gen/kitex_gen/cart"
	"gomall/rpc_gen/kitex_gen/checkout"
	order "gomall/rpc_gen/kitex_gen/order"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
	if req.UserId <= 0 {
		return nil, kerrors.NewGRPCBizStatusError(6004002, "用户id无效！")
	}
	var respOrders []*order.Order
	orders, err := model.ListOrder(s.ctx, mysql.DB, req.GetUserId())
	for _, o := range orders {
		if o.OrderItems == nil {
			continue
		}
		var orderItems []*order.OrderItem
		for _, item := range o.OrderItems {
			orderItems = append(orderItems, &order.OrderItem{
				CartItem: &cart.CartItem{
					ProductId: item.ProductId,
					Num:       uint32(item.Quantity),
				},
				Cost: item.Cost,
			})
		}
		respOrders = append(respOrders, &order.Order{
			OrderItems:   orderItems,
			OrderId:      o.OrderId,
			UserId:       req.UserId,
			UserCurrency: o.UserCurrency,
			Address: &checkout.Address{
				Street:  o.Address.Street,
				City:    o.Address.City,
				State:   o.Address.State,
				Country: o.Address.Country,
				ZipCode: o.Address.ZipCode,
			},
			Email:     o.Email,
			CreatedAt: o.CreatedAt.Unix(),
		})
	}

	return &order.ListOrderResp{
		Orders: respOrders,
	}, nil
}
